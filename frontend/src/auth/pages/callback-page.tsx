import { useEffect, useState } from 'react';
import { useAuth } from '@/auth/context/auth-context';
import { useNavigate, useSearchParams } from 'react-router-dom';

/**
 * Callback page for OAuth authentication redirects.
 * This component handles the authentication flow after a user signs in with a third-party provider via your Go backend.
 */
export function CallbackPage() {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [error, setError] = useState<string | null>(null);
  const { saveAuth } = useAuth();

  useEffect(() => {
    // Tangkap parameter error jika proses OAuth di backend gagal
    const errorParam = searchParams.get('error');
    const errorDescription = searchParams.get('error_description');

    if (errorParam) {
      setError(errorDescription || 'Authentication failed');
      // Redirect kembali ke halaman signin setelah beberapa detik
      setTimeout(() => {
        navigate(
          `/auth/signin?error=${errorParam}&error_description=${encodeURIComponent(errorDescription || 'Authentication failed')}`,
        );
      }, 1500);
      return;
    }

    const handleCallback = async () => {
      try {
        console.log('Processing OAuth callback from custom backend');

        // Tangkap token yang dikirimkan oleh backend Go melalui URL parameter
        // Biasanya backend akan me-redirect ke: /auth/callback?access_token=xyz...
        const accessToken = searchParams.get('access_token') || searchParams.get('token');
        const refreshToken = searchParams.get('refresh_token') || '';

        if (!accessToken) {
          throw new Error('Token akses tidak ditemukan di URL');
        }

        console.log('Token obtained successfully from URL');

        // Bentuk model auth sesuai standar aplikasi Anda
        const authModel = {
          access_token: accessToken,
          refresh_token: refreshToken,
        };

        // Simpan token ke localStorage secara manual agar BackendAdapter bisa membacanya
        localStorage.setItem('access_token', accessToken);

        // Simpan data auth ke dalam React Context
        saveAuth(authModel);
        console.log('Auth data saved to context');

        // Ambil URL tujuan berikutnya (jika ada) atau arahkan ke root (dashboard)
        const nextPath = searchParams.get('next') || '/';

        console.log('Redirecting to:', nextPath);
        navigate(nextPath);
      } catch (err) {
        console.error('Error processing OAuth callback:', err);
        setError('An unexpected error occurred during authentication');

        setTimeout(() => {
          navigate(
            '/auth/signin?error=auth_callback_error&error_description=Failed to complete authentication',
          );
        }, 1500);
      }
    };

    handleCallback();
  }, [navigate, searchParams, saveAuth]);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen p-4 text-center">
      {error ? (
        <div className="space-y-4">
          <h2 className="text-xl font-semibold text-destructive">
            Authentication Error
          </h2>
          <p className="text-muted-foreground">{error}</p>
          <p className="text-sm">Redirecting to sign-in page...</p>
        </div>
      ) : (
        <div className="space-y-4">
          {/* Anda bisa menambahkan animasi loading di sini jika mau */}
          <h2 className="text-xl font-semibold">Authenticating...</h2>
          <p className="text-muted-foreground">Please wait while we complete your sign in.</p>
        </div>
      )}
    </div>
  );
}