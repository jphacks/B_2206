/** @type {import('next').NextConfig} */
const isProd = process.env.NEXT_PUBLIC_APP_ENV === 'production'

const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  env: {
    SSR_API_URI: isProd
      ? 'https://sumimatch-api.nutfes.net'
      : 'http://nutfes-finansu-api:1323',
    CSR_API_URI: isProd
      ? 'https://sumimatch-api.nutfes.net'
      : 'http://localhost:1323',
    CLIENT_URL: isProd
      ? 'https://sumimatch.nutfes.net'
      : 'http://localhost:3000',
  },
}

module.exports = nextConfig
