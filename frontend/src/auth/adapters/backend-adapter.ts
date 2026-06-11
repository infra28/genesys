import { AuthModel, UserModel } from '@/auth/lib/models';

// Gunakan URL dari file .env Anda
const API_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api';

/**
 * Helper function untuk mendapatkan header otorisasi
 */
const getAuthHeaders = (): HeadersInit => {
  const token = localStorage.getItem('access_token');
  return {
    'Content-Type': 'application/json',
    ...(token ? { 'Authorization': `Bearer ${token}` } : {}),
  };
};

/**
 * Backend adapter that maintains the same interface as the existing auth flow
 * but uses your custom Go API under the hood.
 */
export const BackendAdapter = {
  /**
   * Login with email and password
   */
  async login(email: string, password: string): Promise<AuthModel> {
    console.log('BackendAdapter: Attempting login with email:', email);

    try {
      const response = await fetch(`${API_URL}/auth/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        console.error('Login error:', errorData);
        throw new Error(errorData.error || 'Login gagal. Periksa email dan password Anda.');
      }

      const data = await response.json();

      // Simpan token untuk digunakan pada request selanjutnya
      if (data.access_token) {
        localStorage.setItem('access_token', data.access_token);
      }

      return {
        access_token: data.access_token,
        refresh_token: data.refresh_token,
      };
    } catch (error) {
      console.error('BackendAdapter: Unexpected login error:', error);
      throw error;
    }
  },

  /**
   * Login with OAuth provider (Google, GitHub, etc.)
   */
  async signInWithOAuth(
    provider: string,
    options?: { redirectTo?: string }
  ): Promise<void> {
    console.log(`BackendAdapter: Initiating OAuth flow for ${provider}`);
    // Catatan: Anda perlu mengimplementasikan endpoint OAuth di backend Go Anda.
    // Biasanya ini dilakukan dengan me-redirect user ke endpoint backend,
    // yang kemudian akan me-redirect ke Google/Github.
    const redirectTo = options?.redirectTo || `${window.location.origin}/auth/callback`;
    window.location.href = `${API_URL}/auth/${provider}?redirect_to=${encodeURIComponent(redirectTo)}`;
  },

  /**
   * Register a new user
   */
  async register(
    email: string,
    password: string,
    password_confirmation: string,
    firstName?: string,
    lastName?: string,
  ): Promise<AuthModel> {
    if (password !== password_confirmation) {
      throw new Error('Passwords do not match');
    }

    const response = await fetch(`${API_URL}/auth/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email,
        password,
        first_name: firstName || '',
        last_name: lastName || '',
        // Format username otomatis seperti logika Supabase sebelumnya
        username: email.split('@')[0],
      }),
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || 'Gagal mendaftarkan akun');
    }

    const data = await response.json();

    // Jika registrasi langsung memberikan token (tidak perlu verifikasi email)
    if (data.access_token) {
      localStorage.setItem('access_token', data.access_token);
      return {
        access_token: data.access_token,
        refresh_token: data.refresh_token,
      };
    }

    return { access_token: '', refresh_token: '' };
  },

  /**
   * Request password reset
   */
  async requestPasswordReset(email: string): Promise<void> {
    const response = await fetch(`${API_URL}/auth/forgot-password`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || 'Gagal meminta reset password');
    }
  },

  /**
   * Reset password with token
   */
  async resetPassword(
    password: string,
    password_confirmation: string,
  ): Promise<void> {
    if (password !== password_confirmation) {
      throw new Error('Passwords do not match');
    }

    // Biasanya endpoint ini membutuhkan token dari parameter URL
    // Di sini kita asumsikan token dikirim via header jika user sedang login,
    // atau Anda mungkin perlu memodifikasi ini untuk menerima token dari URL.
    const response = await fetch(`${API_URL}/auth/reset-password`, {
      method: 'POST',
      headers: getAuthHeaders(),
      body: JSON.stringify({ password }),
    });

    if (!response.ok) throw new Error('Gagal mereset password');
  },

  /**
   * Request another verification email
   */
  async resendVerificationEmail(email: string): Promise<void> {
    const response = await fetch(`${API_URL}/auth/resend-verification`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email }),
    });

    if (!response.ok) throw new Error('Gagal mengirim ulang email verifikasi');
  },

  /**
   * Get current user from the session
   */
  async getCurrentUser(): Promise<UserModel | null> {
    const token = localStorage.getItem('access_token');
    if (!token) return null;

    try {
      return await this.getUserProfile();
    } catch (error) {
      // Jika token expired atau invalid, bersihkan token
      localStorage.removeItem('access_token');
      return null;
    }
  },

  /**
   * Get user profile
   */
  async getUserProfile(): Promise<UserModel> {
    const response = await fetch(`${API_URL}/users/me`, {
      method: 'GET',
      headers: getAuthHeaders(),
    });

    if (!response.ok) {
      throw new Error('Gagal mengambil profil user. Sesi mungkin telah berakhir.');
    }

    const userData = await response.json();

    // Pastikan mapping data dari backend Go sesuai dengan UserModel di frontend
    return {
      id: userData.id,
      email: userData.email || '',
      email_verified: userData.email_verified || false,
      username: userData.username || '',
      first_name: userData.first_name || '',
      last_name: userData.last_name || '',
      fullname: userData.fullname || `${userData.first_name || ''} ${userData.last_name || ''}`.trim(),
      occupation: userData.occupation || '',
      company_name: userData.company_name || '',
      phone: userData.phone || '',
      roles: userData.roles || [],
      pic: userData.pic || '',
      language: userData.language || 'en',
      is_admin: userData.is_admin || false,
    };
  },

  /**
   * Update user profile
   */
  async updateUserProfile(userData: Partial<UserModel>): Promise<UserModel> {
    const response = await fetch(`${API_URL}/users/me`, {
      method: 'PUT', // atau PATCH, sesuaikan dengan backend Go Anda
      headers: getAuthHeaders(),
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      throw new Error(errorData.message || 'Gagal memperbarui profil');
    }

    // Ambil data profil terbaru setelah di-update
    return this.getUserProfile();
  },

  /**
   * Logout the current user
   */
  async logout(): Promise<void> {
    try {
      // Panggil endpoint logout di backend (opsional, untuk invalidasi token/sesi di server)
      await fetch(`${API_URL}/auth/logout`, {
        method: 'POST',
        headers: getAuthHeaders(),
        credentials: 'include',
      });
    } catch (error) {
      console.warn('Gagal menghubungi server saat logout, tetap membersihkan sesi lokal.');
    } finally {
      // Yang paling penting: bersihkan token dari localStorage
      localStorage.removeItem('access_token');
    }
  },
};