import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: true,
  // Proxy API requests to the Go backend during development.
  // In production, set NEXT_PUBLIC_API_BASE_URL to the Go server's address
  // and these rewrites become a no-op since the frontend calls the URL directly.
  async rewrites() {
    return [
      {
        source: "/video-info",
        destination: process.env.BACKEND_URL
          ? `${process.env.BACKEND_URL}/video-info`
          : "http://localhost:8080/video-info",
      },
      {
        source: "/download",
        destination: process.env.BACKEND_URL
          ? `${process.env.BACKEND_URL}/download`
          : "http://localhost:8080/download",
      },
    ];
  },
};

export default nextConfig;
