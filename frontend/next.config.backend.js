import { nextConfig } from "./next.config";

const backendConfig = {
  ...nextConfig,
  // Ensure backend only builds the API routes and necessary server components
  output: 'standalone',
  // Optional: Add environment variables specific to backend deployment
  env: {
    ...nextConfig.env,
    BACKEND_URL: 'https://140.238.71.103',
  }
};

export default backendConfig; 