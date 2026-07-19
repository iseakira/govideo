const nextConfig = {
  env: {
    NEXT_PUBLIC_AUTH_ENABLE: process.env.AUTH_ENABLE ?? 'false',
    NEXT_PUBLIC_APP_URL: process.env.APP_URL ?? 'http://localhost:3000',
  },
  async rewrites() {
    return process.env.NODE_ENV === 'production'
      ? [
          {
            source: '/api/v1/stream/:path*',
            destination: `${process.env.STREAM_API_URL}/api/v1/stream/:path*`,
          },
        ]
      : [];
  },
};

export default nextConfig;
