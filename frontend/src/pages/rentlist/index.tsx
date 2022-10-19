import type { NextPage } from 'next'
import { Card, Label } from '@components/common'
import clsx from 'clsx'
import { User, Area, Value, Tag } from '@type/common'

const Rentlist: NextPage = () => {
  const user: User = {
    id: 1,
    name: 'test',
    email: 'test',
    password: 'test',
    personalInfoId: 1,
    companyInfoId: 1,
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
  }
  console.log(data)
  const dataList = [data, data, data, data]
  // const displayData = [{ user: user, request: request, detail: detail }]

  return (
    <>
      <div className="px-10 pt-10 pb-3">
        <p className="text-start border-primary-1 border-l-8 pl-2 text-4xl">
          {data.request.detail.area.prefecture} {data.request.detail.area.city}
          で物件を探している人たち
        </p>
      </div>
      <div className="grid w-full sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
        {dataList.map((data) => (
          <Card width="w-1/1">
            {/* <div className="py-2">
              <p className="border-primary-1 border-b-2 text-left text-xl font-bold">
                立地
              </p>
            </div>
            <div className="text-left">
              <p>
                {data.request.detail.area.prefecture}
                {data.request.detail.area.city
                  ? data.request.detail.area.city
                  : ''}
                {data.request.detail.area.addressNumber
                  ? data.request.detail.area.addressNumber
                  : ''}
              </p>
            </div> */}
            <div className="py-2">
              <p className="border-primary-1 border-b-2 text-left text-xl font-bold">
                希望する条件
              </p>
            </div>
            <div className="grid grid-cols-2 text-left">
              {data.request.detail.classification.range.map((range) => (
                <>
                  {range.name === '築年数' ||
                    (range.name === '賃料' && (
                      <p className="py-2 text-lg">
                        {range.name}
                        <p className="text-sm">
                          {range.values.map((value, index) => (
                            <span>
                              {value.name}
                              {index < range.values.length - 1 && '〜'}
                            </span>
                          ))}
                        </p>
                      </p>
                    ))}
                </>
              ))}
              {data.request.detail.classification.value.map((valuelist) => (
                <p className="text-lg">
                  {valuelist.name}
                  <p className="text-sm">
                    {valuelist.values.map((value) => (
                      <span>{value.name}</span>
                    ))}
                  </p>
                </p>
              ))}
            </div>
            {/* <div className="py-2">
              <p className="border-primary-1 border-b-2 text-left text-xl font-bold">
                より詳しい条件
              </p>
            </div> */}
            <div className="text-left">
              {data.request.detail.classification.tag.map((taglist) => (
                <>
                  {taglist.name === '間取り' && (
                    <div className="py-2 text-lg">
                      {taglist.name}
                      <div className="flex text-sm">
                        {taglist.tags.map((tag) => (
                          <Label
                            name={tag.name}
                            width={'w-auto'}
                            className={'mx-1'}
                          />
                        ))}
                      </div>
                    </div>
                  )}
                </>
              ))}
            </div>
          </Card>
        ))}
      </div>
    </>
  )
}

export default Rentlist
