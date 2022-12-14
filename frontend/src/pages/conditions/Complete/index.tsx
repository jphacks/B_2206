import Link from 'next/link'
import { useEffect } from 'react'
import { useRecoilState } from 'recoil'
import { Card, PrimaryButton } from '@components/common'
import { conditionState } from '@components/store/Condition/condition'

function Complete(): JSX.Element {
  const [condition, setCondition] = useRecoilState(conditionState)

  useEffect(() => {
    scrollTo(0, 0)
  }, [])

  return (
    <div className="w-full">
      <Card width={'w-4/5'}>
        <div className={'text-primary-2 flex justify-center text-xl'}>
          <p>登録が完了しました!</p>
        </div>
        <div className="inline-flex w-full items-center justify-center">
          <hr className="w-9/10 bg-primary-1 my-8 h-px border-0" />
          <span className="absolute left-1/2 -translate-x-1/2 bg-white px-3 font-medium text-gray-900 dark:bg-gray-900 dark:text-white">
            詳細
          </span>
        </div>
        <div className="w-9/10 mx-auto">
          <div className="flex flex-col gap-3">
            <div className="flex gap-3">
              <p>賃料</p>
              <p>{condition.conditions.priceRange}</p>
            </div>
            <div className="mx-3 flex flex-row gap-3">
              <p>賃料の条件</p>
              {condition.conditions.priceOptions.map((option: string) => {
                return <p key={option}>{option}</p>
              })}
            </div>
            <div className="flex flex-row gap-3">
              <p>間取りタイプ</p>
              {condition.conditions.roomTypes.map((roomType: string) => {
                return <p key={roomType}>{roomType}</p>
              })}
            </div>
            <div className="flex flex-row gap-3">
              <p>建物の種類</p>
              {condition.conditions.buildTypes.map((buildType: string) => {
                return <p key={buildType}>{buildType}</p>
              })}
            </div>
            <div className="flex flex-row gap-3">
              <p>駅徒歩</p>
              <p>{condition.conditions.time}</p>
            </div>
            <div className="flex flex-row gap-3">
              <p>占有面積</p>
              <p>{condition.conditions.areaRange}</p>
            </div>
            <div className="flex flex-row gap-3">
              <p>築年数</p>
              <p>{condition.conditions.buildAge}</p>
            </div>
            <div className="flex flex-row gap-3">
              <p>こだわり条件</p>
              {condition.conditions.otherConditions.map((other: string) => {
                return <p key={other}>{other}</p>
              })}
            </div>
          </div>
        </div>
        <div className="inline-flex w-full items-center justify-center">
          <hr className="w-9/10 bg-primary-1 my-8 h-px border-0" />
        </div>
        <div className={'mt-10 mb-5 flex justify-center'}>
          <Link href={'/'}>
            <PrimaryButton>トップページに戻る</PrimaryButton>
          </Link>
        </div>
      </Card>
    </div>
  )
}

export default Complete
