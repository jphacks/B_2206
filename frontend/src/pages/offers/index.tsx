import clsx from 'clsx'
import type { NextPage } from 'next'
import { useState, useEffect, memo, useRef } from 'react'
import { Card, Label, Modal, PrimaryButton, Tooltip } from '@components/common'
import { Close, Domain, Phone, WebSite } from '@components/icons'
import {
  User,
  Area,
  Value,
  Tag,
  PersonalInfo,
  CompanyInfo,
  Matching,
} from '@type/common'

interface Props {
  children?: React.ReactNode
  data: any
}

interface CardContentProps {
  children?: React.ReactNode
  userData: {
    user: User
    personalInfo: PersonalInfo
    companyInfo: CompanyInfo
    matching: Matching
  }
}

const Rentlist: NextPage = () => {
  const personalInfo: PersonalInfo = {
    id: 1,
    familyName: '藤崎',
    firstName: '竜成',
    familyNameKana: 'フジサキ',
    firstNameKana: 'リュウセイ',
    birthDay: '19990202',
    phoneNumber: '09012345678',
    userId: 1,
    created_at: '',
    updated_at: '',
  }
  const companyInfo: CompanyInfo = {
    id: 1,
    name: '長岡金型',
    phoneNumber: '11112345678',
    postCode: '9456666',
    prefecture: '新潟県',
    city: '長岡市西陵町',
    addressNumber: '2674-31',
    buildingName: '長岡金型ビル',
    website: 'https://nagaoka-kanagata.co.jp',
    userId: 1,
    created_at: '',
    updated_at: '',
  }
  const matching: Matching = {
    id: 1,
    requestId: 1,
    buyerUserId: 1,
    sellerUserId: 2,
    status: 'オファー中',
    sellerMessage:
      '新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。',
    created_at: '',
    updated_at: '',
  }
  const user: User = {
    id: 1,
    name: 'test',
    email: 'test',
    password: 'test',
    created_at: '',
    updated_at: '',
  }
  const area: Area = {
    id: 1,
    postCode: 1110000,
    prefecture: '北海道',
    city: '札幌市',
    addressNumber: '1-1-1',
    buildingName: 'test',
    created_at: '',
    updated_at: '',
  }

  const tagList: Tag[] = [
    { id: 1, name: 'IHコンロ', created_at: '', updated_at: '' },
    { id: 2, name: 'バス・トイレ別', created_at: '', updated_at: '' },
    { id: 3, name: '1R', created_at: '', updated_at: '' },
    { id: 4, name: '1K', created_at: '', updated_at: '' },
  ]
  const tags = [
    {
      name: '条件',
      tags: [tagList[0], tagList[1]],
    },
    {
      name: '間取り',
      tags: [tagList[2], tagList[3]],
    },
  ]
  const valueList: Value[] = [
    {
      id: 1,
      name: '20㎡',
      value: 20,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '30㎡',
      value: 30,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '3万円',
      value: 3,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '4万円',
      value: 4,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: '新築',
      value: 0,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
  ]
  const values = [
    {
      name: '築年数',
      values: [valueList[4]],
    },
  ]
  const ranges = [
    {
      name: '面積',
      values: [valueList[0], valueList[1]],
    },
    {
      name: '賃料',
      values: [valueList[2], valueList[3]],
    },
  ]
  const classifications = {
    value: values,
    range: ranges,
    tag: tags,
  }
  const details = {
    area: area,
    classification: classifications,
  }
  const requests = {
    detail: details,
  }
  const data = {
    user: user,
    request: requests,
    matching: matching,
  }
  const dataList = [data, data, data, data]
  // const displayData = [{ user: user, request: request, detail: detail }]

  const userData = {
    user: user,
    personalInfo: personalInfo,
    companyInfo: companyInfo,
    matching: matching,
  }
  const userDataList = [userData, userData, userData, userData, userData]

  // 検索条件を表示する用のstate
  const [isConditionOpen, setIsConditionOpen] = useState<boolean>(false)
  const onConditionOpen = () => setIsConditionOpen(true)
  const onConditionClose = () => setIsConditionOpen(false)

  // メッセージを送る用のstate
  const [isMessageOpen, setIsMessageOpen] = useState<boolean>(false)
  const onMessageOpen = () => setIsMessageOpen(true)
  const onMessageClose = () => setIsMessageOpen(false)
  // カードの中に表示するデータ
  const CardContent: React.FC<CardContentProps> = memo((props) => {
    return (
      <>
        <div className="mt-2 text-left">
          <div className="items-center justify-start">
            <div className="flex">
              <div className="flex items-center">
                <Domain width="24" height="24" />
                <span className="py-1 pl-1 pl-1 text-left text-xl font-bold">
                  {props.userData.companyInfo.name}
                </span>
              </div>
              <div className="mx-auto"></div>
              <div className="flex items-center">
                <div className="pr-2">
                  <Tooltip content={props.userData.companyInfo.phoneNumber}>
                    <a href={`tel:${props.userData.companyInfo.phoneNumber}`}>
                      <Phone width="24" height="24" />
                    </a>
                  </Tooltip>
                </div>
                <Tooltip content={props.userData.companyInfo.website}>
                  <a href={props.userData.companyInfo.website} target="_blank" rel="noreferrer">
                    <WebSite width="24" height="24" />
                  </a>
                </Tooltip>
              </div>
            </div>
          </div>
          <div className="border-primary-2 flex items-center justify-start border-t-2 border-dotted pt-2">
            <span className="text-md font-bold">
              {props.userData.matching.sellerMessage}
            </span>
          </div>
          <div className="mt-4 flex items-center justify-center">
            <PrimaryButton
              className="sm:text-lg lg:text-xl"
              onClick={onMessageOpen}
            >
              オファーの詳細を確認
            </PrimaryButton>
          </div>
        </div>
      </>
    )
  })

  // 条件確認のモーダルに表示するデータ
  const ConditionModalContent: React.FC<Props> = memo((props) => {
    return (
      <div>
        <Modal width="w-auto">
          <div
            className={clsx(
              'justify-left ml-4 mb-4 w-fit cursor-pointer items-center',
            )}
          >
            <Close width="32" height="32" onClick={onConditionClose} />
          </div>
          <div>
            <span className="border-primary-1 border-b-2 text-left text-xl font-bold">
              希望する条件
            </span>
          </div>
          <div className="mt-1 grid grid-cols-3 gap-8 text-left">
            {props.data.request.detail.classification.range.map(
              (range: any) => (
                <span className="pt-2 pb-1 text-lg font-bold">
                  {range.name}
                  <br />
                  <span className="text-sm">
                    {range.values.map((value: Value, index: number) => (
                      <span>
                        {value.name}
                        {index < range.values.length - 1 && '〜'}
                      </span>
                    ))}
                  </span>
                </span>
              ),
            )}
            {props.data.request.detail.classification.value.map(
              (valuelist: any) => (
                <>
                  {valuelist.name === '築年数' && (
                    <span className="pt-2 pb-1 text-lg font-bold">
                      {valuelist.name}
                      <br />
                      <span className="text-sm">
                        {valuelist.values.map((value: Value) => (
                          <span>{value.name}</span>
                        ))}
                      </span>
                    </span>
                  )}
                </>
              ),
            )}
          </div>
          <div className="text-left">
            {props.data.request.detail.classification.tag.map(
              (taglist: any) => (
                <div className="py-1 text-lg font-bold">
                  {taglist.name}
                  <br />
                  <div className="w-1/1 flex flex-wrap text-sm">
                    {taglist.tags.map((tag: Tag) => (
                      <Label
                        name={tag.name}
                        width={'w-auto'}
                        className={'m-1'}
                      />
                    ))}
                  </div>
                </div>
              ),
            )}
          </div>
        </Modal>
      </div>
    )
  })

  // オファー確認のモーダルに表示するデータ
  const MessageModalContent: React.FC<Props> = memo((props) => {
    return (
      <div>
        <Modal width="w-auto">
          <div
            className={clsx(
              'justify-left ml-4 mb-4 w-fit cursor-pointer items-center',
            )}
          >
            <Close width="32" height="32" onClick={onMessageClose} />
          </div>
          <div>
            <span className="text-left text-xl font-bold">希望する条件</span>
          </div>
          <div className="border-primary-1 grid grid-cols-3 gap-8 border-t-2 border-dotted text-left ">
            {props.data.request.detail.classification.range.map(
              (range: any) => (
                <span className="pt-2 pb-1 text-lg font-bold">
                  {range.name}
                  <br />
                  <span className="text-sm">
                    {range.values.map((value: Value, index: number) => (
                      <span>
                        {value.name}
                        {index < range.values.length - 1 && '〜'}
                      </span>
                    ))}
                  </span>
                </span>
              ),
            )}
            {props.data.request.detail.classification.value.map(
              (valuelist: any) => (
                <>
                  {valuelist.name === '築年数' && (
                    <span className="pt-2 pb-1 text-lg font-bold">
                      {valuelist.name}
                      <br />
                      <span className="text-sm">
                        {valuelist.values.map((value: Value) => (
                          <span>{value.name}</span>
                        ))}
                      </span>
                    </span>
                  )}
                </>
              ),
            )}
          </div>
          <div className="mb-3 grid grid-cols-2 gap-3 text-left">
            {props.data.request.detail.classification.tag.map(
              (taglist: any) => (
                <div className="col-span-1 py-1 text-lg font-bold">
                  {taglist.name}
                  <br />
                  <div className="w-1/1 flex flex-wrap text-sm">
                    {taglist.tags.map((tag: Tag) => (
                      <Label
                        name={tag.name}
                        width={'w-auto'}
                        className={'m-1'}
                      />
                    ))}
                  </div>
                </div>
              ),
            )}
          </div>
          <div>
            <span className="mt-3 text-left text-xl font-bold">
              希望する条件
            </span>
          </div>
          <div className="border-primary-2 flex items-center justify-start border-t-2 border-dotted pt-2">
            <span className="text-md font-bold">
              {props.data.matching.sellerMessage}
            </span>
          </div>
          <div className="flex items-center justify-center pt-8">
            <PrimaryButton
              className="sm:text-lg lg:text-xl"
              onClick={onMessageClose}
            >
              オファーを承諾
            </PrimaryButton>
          </div>
        </Modal>
      </div>
    )
  })

  return (
    <>
      <div className="flex px-10 pt-10 pb-3">
        <span className="text-start border-primary-1 border-l-8 pl-2 sm:text-xl lg:text-4xl">
          オファー一覧
        </span>
        <div className="mx-auto"></div>
        <PrimaryButton
          className="sm:text-lg lg:text-2xl"
          onClick={onConditionOpen}
        >
          検索条件を確認
        </PrimaryButton>
      </div>
      {isConditionOpen && <ConditionModalContent data={data} />}
      <div className="w-9/10 mx-auto grid sm:grid-cols-1 lg:grid-cols-2">
        {userDataList.map((userData) => (
          <Card width="w-1/1" m="m-4">
            <CardContent userData={userData}></CardContent>
          </Card>
        ))}
        {isMessageOpen && <MessageModalContent data={data} />}
      </div>
    </>
  )
}

export default Rentlist
