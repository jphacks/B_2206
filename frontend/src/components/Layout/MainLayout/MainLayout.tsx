import clsx from 'clsx'
import Head from 'next/head'
import React from 'react'

import s from './MainLayout.module.css'
import Header from '@components/Layout/Header'

interface User {
  id: number
  name: string
  department_id: number
  role_id: number
}

interface LayoutProps {
  children?: React.ReactNode
  currentUser?: User
}

export default function MainLayout(props: LayoutProps) {
  return (
    <>
      <Head>
        <title>SumiMatch</title>
        <meta name="" content="" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className={clsx('h-screen w-full')}>
        <div className={clsx('h-16 w-full')}>
          <Header />
        </div>
        <div className={clsx(s.parent)}>
          <div className={clsx('h-full w-full', s.content)}>
            {props.children}
          </div>
        </div>
      </div>
    </>
  )
}
