import { EventEmitter } from 'stream'
import clsx from 'clsx'
import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import Link from 'next/link'
import React, { useEffect, useCallback, useState } from 'react'
import { useRecoilState } from 'recoil'
import styles from '@styles/Home.module.css'
import { Card, PrimaryButton, Radio } from '@components/common'
import { conditionState } from '@components/store/Condition/condition'

import City from './City'
import Complete from './Complete'
import Other from './Other'
import Prefecture from './Prefecuture'
import { userState } from '@components/store/Auth/auth'

const Modals = {
  prefecture: 'prefecture',
  city: 'city',
  other: 'other',
  complete: 'complete',
}

const Home: NextPage = () => {
  const [condition, setConditon] = useRecoilState(conditionState)
  const [modalName, setModalName] = useState(Modals.prefecture)

  const handleClickClose = useCallback(() => {
    setModalName(Modals.prefecture)
    document.removeEventListener('click', handleClickClose)
  }, [])

  useEffect(() => {
    return () => {
      document.removeEventListener('click', handleClickClose)
    }
  }, [handleClickClose])

  const PrefectureModal = (
    <Prefecture
      setModalName={setModalName}
      nextModalName={Modals.city}
      otherModalName={Modals.other}
    />
  )

  const CityModal = (
    <City
      setModalName={setModalName}
      prevModalName={Modals.prefecture}
      nextModalName={Modals.other}
    />
  )

  const OtherModal = (
    <Other
      setModalName={setModalName}
      prevModalName={Modals.city}
      morePrevModalName={Modals.prefecture}
      nextModalName={Modals.complete}
    />
  )

  const CompleteModal = <Complete />

  return (
    <div>
      {modalName === Modals.prefecture && PrefectureModal}
      {modalName === Modals.city && CityModal}
      {modalName === Modals.other && OtherModal}
      {modalName === Modals.complete && CompleteModal}
    </div>
  )
}

export default Home
