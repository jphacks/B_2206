import type { NextPage } from 'next'
import Link from 'next/link'
import Head from 'next/head'
import Image from 'next/image'
import styles from '@styles/Home.module.css'
import {
  Card,
  PrimaryButton,
  Radio,
} from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import clsx from 'clsx'
import React,{ useEffect,useCallback,useState }  from 'react';
import { EventEmitter } from 'stream'
import Prefecture from './Prefecuture'
import City from './city'

const Modals = {
    prefecture: "prefecture",
    city: "city",
}

const Home: NextPage = () => {
    const [user, setUser] = useRecoilState(userState)
    const [modalName, setModalName] = useState(Modals.prefecture);
    // console.log(user)

    const handleClickClose = useCallback(() => {
        setModalName(Modals.prefecture);
        document.removeEventListener('click', handleClickClose)
    },[])

    useEffect(() => {
        return () => {
            document.removeEventListener('click', handleClickClose)
        }
    },[handleClickClose])
  
    const PrefectureModal = (
        <Prefecture nextSetModalName={setModalName} nextModalName={Modals.city} />
    )

    const CityModal = (
        <City prevSetModalName={setModalName} prevModalName={Modals.prefecture} />
    )

    return (
      <div>
        {modalName === Modals.prefecture && PrefectureModal}
        {modalName === Modals.city && CityModal}
      </div>
    )
  }
  
  export default Home
  