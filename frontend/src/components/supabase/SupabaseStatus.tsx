import React, { useEffect, useState } from 'react';

/**
 * A simple component that displays the status of the Backend API connection.
 * This can be used during development to verify that your Go backend is running and accessible.
 */
export const ApiStatus: React.FC = () => {
  const [status, setStatus] = useState<'checking' | 'connected' | 'error'>(
    'checking',
  );
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const checkConnection = async () => {
      try {
        const apiUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';
        
        // Memanggil endpoint publik di backend Go Anda. 
        // Sangat disarankan untuk membuat endpoint GET /health di backend Go yang hanya mereturn HTTP 200 OK
        const response = await fetch(`${apiUrl}/health`);
        
        if (response.ok) {
          setStatus('connected');
        } else {
          setStatus('error');
          setError(`API returned status: ${response.status}`);
        }
      } catch (e) {
        setStatus('error');
        setError(e instanceof Error ? e.message : 'Network error or server is down');
      }
    };

    checkConnection();
  }, []);

  return (
    <div className="p-4 rounded-md border">
      <h3 className="text-lg font-medium mb-2">Backend API Status</h3>
      <div className="flex items-center gap-2">
        <div
          className={`w-3 h-3 rounded-full ${
            status === 'checking'
              ? 'bg-accent'
              : status === 'connected'
                ? 'bg-green-500'
                : 'bg-red-500'
          }`}
        />
        <span>
          {status === 'checking'
            ? 'Checking connection...'
            : status === 'connected'
              ? 'Connected to Backend'
              : 'Connection error'}
        </span>
      </div>
      {error && <p className="text-red-500 text-sm mt-2">{error}</p>}
    </div>
  );
};