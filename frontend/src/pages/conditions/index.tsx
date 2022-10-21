import type { NextPage } from 'next'
import Link from 'next/link'
import Head from 'next/head'
import Image from 'next/image'
import styles from '@styles/Home.module.css'
import { Card, PrimaryButton, Radio } from '@components/common'
import { useRecoilState } from 'recoil'
import { conditionState } from '@components/store/Condition/condition'
import clsx from 'clsx'
import React, { useEffect, useCallback, useState } from 'react'
import Prefecture from './Prefecuture'
import City from './City'
import Other from './Other'
import Complete from './Complete'

const Modals = {
  prefecture: 'prefecture',
  city: 'city',
  other: 'other',
  complete: 'complete',
}

const Home: NextPage = () => {
  const [condition, setConditon] = useRecoilState(conditionState)
  const [modalName, setModalName] = useState(Modals.prefecture)

  // console.log(condition)

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
    <Prefecture setModalName={setModalName} nextModalName={Modals.city} />
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
