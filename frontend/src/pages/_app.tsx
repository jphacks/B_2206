import '../styles/globals.css'
import { RecoilRoot } from 'recoil'
import type { AppProps } from 'next/app'
import MainLayout from '@components/Layout/MainLayout'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <RecoilRoot>
      <MainLayout>
        <Component {...pageProps} />
      </MainLayout>
    </RecoilRoot>
  )
}

export default MyApp
