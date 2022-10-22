import { atom } from 'recoil'
import { recoilPersist } from 'recoil-persist'

const { persistAtom } = recoilPersist()

const seedData = {
  // 賃料関係
  rentalPrice: [
    {
      id: 1,
      name: '管理費・共益費込み',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '駐車場代込み',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '礼金なし',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '敷金・保証金なし',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: '初期費用カード決済可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 6,
      name: '家賃カード決済可',
      created_at: '',
      updated_at: '',
    },
  ],
  // 構造
  structure: [
    {
      id: 1,
      name: '鉄筋系',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '鉄骨系',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '木造',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: 'ブロック・その他',
      created_at: '',
      updated_at: '',
    },
  ],
  // 位置
  position: [
    {
      id: 1,
      name: '1階の物件',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '2階以上',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '最上階',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '角部屋',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: '南向き',
      created_at: '',
      updated_at: '',
    },
  ],
  // 室内設備
  roomEquipment: [
    {
      id: 1,
      name: '室内洗濯機置場',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '洗面所独立',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: 'フローリング',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: 'メゾネット',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: 'ロフト',
      created_at: '',
      updated_at: '',
    },
    {
      id: 6,
      name: '防音室',
      created_at: '',
      updated_at: '',
    },
    {
      id: 7,
      name: '地下室',
      created_at: '',
      updated_at: '',
    },
    {
      id: 8,
      name: '家具家電付き',
      created_at: '',
      updated_at: '',
    },
  ],
  // 冷暖房
  airConditioning: [
    {
      id: 1,
      name: 'エアコン付き',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '床暖房',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '灯油暖房',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: 'ガス暖房',
      created_at: '',
      updated_at: '',
    },
  ],
  // バス・トイレ
  bathToilet: [
    {
      id: 1,
      name: 'バス・トイレ別',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '温水洗浄便座',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '浴室乾燥機',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '追い焚き風呂',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: 'シャワールーム',
      created_at: '',
      updated_at: '',
    },
  ],
  // キッチン
  kitchen: [
    {
      id: 1,
      name: 'ガスコンロ対応',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: 'IHコンロ',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: 'コンロ2口以上',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: 'オール電化',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: 'システムキッチン',
      created_at: '',
      updated_at: '',
    },

    {
      id: 6,
      name: 'カウンターキッチン',
      created_at: '',
      updated_at: '',
    },
  ],
  // 建物設備
  buildingEquipment: [
    {
      id: 1,
      name: '駐車場あり',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '駐車場2台以上',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '敷地内駐車場',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '駐輪場あり',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: 'バイク置場あり',
      created_at: '',
      updated_at: '',
    },
    {
      id: 6,
      name: 'エレベーター',
      created_at: '',
      updated_at: '',
    },
    {
      id: 7,
      name: '宅配ボックス',
      created_at: '',
      updated_at: '',
    },
    {
      id: 8,
      name: '敷地内ゴミ箱置場',
      created_at: '',
      updated_at: '',
    },
    {
      id: 9,
      name: 'バルコニー付',
      created_at: '',
      updated_at: '',
    },
    {
      id: 10,
      name: 'ルーフバルコニー付',
      created_at: '',
      updated_at: '',
    },

    {
      id: 11,
      name: '専用庭',
      created_at: '',
      updated_at: '',
    },
    {
      id: 12,
      name: '都市ガス',
      created_at: '',
      updated_at: '',
    },
    {
      id: 13,
      name: 'プロパンガス',
      created_at: '',
      updated_at: '',
    },
    {
      id: 14,
      name: 'バリアフリー',
      created_at: '',
      updated_at: '',
    },
  ],
  // セキュリティ
  security: [
    {
      id: 1,
      name: 'オートロック',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '管理人あり',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: 'TVモニター付インターホン',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '防犯カメラ',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: 'セキュリティ会社加入済',
      created_at: '',
      updated_at: '',
    },
  ],
  // 入居条件
  moveInConditions: [
    {
      id: 1,
      name: '即入居可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '女性限定',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: 'ペット相談可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '楽器相談可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: '事務所利用可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 6,
      name: 'ルームシェア可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 7,
      name: '高齢者歓迎',
      created_at: '',
      updated_at: '',
    },
    {
      id: 8,
      name: 'LGBTフレンドリー',
      created_at: '',
      updated_at: '',
    },
    {
      id: 9,
      name: 'カスタマイズ可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 10,
      name: 'DIY可',
      created_at: '',
      updated_at: '',
    },

    {
      id: 11,
      name: '定期借家を含まない',
      created_at: '',
      updated_at: '',
    },
  ],
  // テレビ・通信
  tvAndCommunication: [
    {
      id: 1,
      name: 'インターネット接続可',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: 'BSアンテナ',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: 'CSアンテナ',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: 'ケーブルテレビ',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: 'インターネット無料',
      created_at: '',
      updated_at: '',
    },
  ],
  // 収納
  storage: [
    {
      id: 1,
      name: '床下収納',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: 'シューズボックス',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: 'トランクルーム',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: 'ウォークインクローゼット',
      created_at: '',
      updated_at: '',
    },
  ],
  // 間取り
  roomLayout: [
    {
      id: 1,
      name: 'ワンルーム',
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '1K',
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '1DK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '1LDK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: '2K',
      created_at: '',
      updated_at: '',
    },
    {
      id: 6,
      name: '2DK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 7,
      name: '2LDK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 8,
      name: '3K',
      created_at: '',
      updated_at: '',
    },
    {
      id: 9,
      name: '3DK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 10,
      name: '3LDK',
      created_at: '',
      updated_at: '',
    },

    {
      id: 11,
      name: '4K',
      created_at: '',
      updated_at: '',
    },
    {
      id: 12,
      name: '4DK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 13,
      name: '4LDK',
      created_at: '',
      updated_at: '',
    },
    {
      id: 14,
      name: '5K以上',
      created_at: '',
      updated_at: '',
    },
  ],
  // 賃料
  price: [
    { id: 1, name: '3万円', value: 3, created_at: '', updated_at: '' },
    { id: 2, name: '3.5万円', value: 3.5, created_at: '', updated_at: '' },
    { id: 3, name: '4万円', value: 4, created_at: '', updated_at: '' },
    { id: 4, name: '4.5万円', value: 4.5, created_at: '', updated_at: '' },
    { id: 5, name: '5万円', value: 5, created_at: '', updated_at: '' },
    { id: 6, name: '5.5万円', value: 5.5, created_at: '', updated_at: '' },
    { id: 7, name: '6万円', value: 6, created_at: '', updated_at: '' },
    { id: 8, name: '6.5万円', value: 6.5, created_at: '', updated_at: '' },
    { id: 9, name: '7万円', value: 7, created_at: '', updated_at: '' },
    { id: 10, name: '7.5万円', value: 7.5, created_at: '', updated_at: '' },
    { id: 11, name: '8万円', value: 8, created_at: '', updated_at: '' },
    { id: 12, name: '8.5万円', value: 8.5, created_at: '', updated_at: '' },
    { id: 13, name: '9万円', value: 9, created_at: '', updated_at: '' },
    { id: 14, name: '9.5万円', value: 9.5, created_at: '', updated_at: '' },
    { id: 15, name: '10万円', value: 10, created_at: '', updated_at: '' },
    { id: 16, name: '10.5万円', value: 10.5, created_at: '', updated_at: '' },
    { id: 17, name: '11万円', value: 11, created_at: '', updated_at: '' },
    { id: 18, name: '11.5万円', value: 11.5, created_at: '', updated_at: '' },
    { id: 19, name: '12万円', value: 12, created_at: '', updated_at: '' },
    { id: 20, name: '12.5万円', value: 12.5, created_at: '', updated_at: '' },
    { id: 21, name: '13万円', value: 13, created_at: '', updated_at: '' },
    { id: 22, name: '13.5万円', value: 13.5, created_at: '', updated_at: '' },
    { id: 23, name: '14万円', value: 14, created_at: '', updated_at: '' },
    { id: 24, name: '14.5万円', value: 14.5, created_at: '', updated_at: '' },
    { id: 25, name: '15万円', value: 15, created_at: '', updated_at: '' },
    { id: 26, name: '15.5万円', value: 15.5, created_at: '', updated_at: '' },
    { id: 27, name: '16万円', value: 16, created_at: '', updated_at: '' },
    { id: 28, name: '16.5万円', value: 16.5, created_at: '', updated_at: '' },
    { id: 29, name: '17万円', value: 17, created_at: '', updated_at: '' },
    { id: 30, name: '17.5万円', value: 17.5, created_at: '', updated_at: '' },
    { id: 31, name: '18万円', value: 18, created_at: '', updated_at: '' },
    { id: 32, name: '18.5万円', value: 18.5, created_at: '', updated_at: '' },
    { id: 33, name: '19万円', value: 19, created_at: '', updated_at: '' },
    { id: 34, name: '19.5万円', value: 19.5, created_at: '', updated_at: '' },
    { id: 35, name: '20万円', value: 20, created_at: '', updated_at: '' },
    { id: 36, name: '21万円', value: 21, created_at: '', updated_at: '' },
    { id: 37, name: '22万円', value: 22, created_at: '', updated_at: '' },
    { id: 38, name: '23万円', value: 23, created_at: '', updated_at: '' },
    { id: 39, name: '24万円', value: 24, created_at: '', updated_at: '' },
    { id: 40, name: '25万円', value: 25, created_at: '', updated_at: '' },
    { id: 41, name: '26万円', value: 26, created_at: '', updated_at: '' },
    { id: 42, name: '27万円', value: 27, created_at: '', updated_at: '' },
    { id: 43, name: '28万円', value: 28, created_at: '', updated_at: '' },
    { id: 44, name: '29万円', value: 29, created_at: '', updated_at: '' },
    { id: 45, name: '30万円', value: 30, created_at: '', updated_at: '' },
    { id: 46, name: '35万円', value: 35, created_at: '', updated_at: '' },
    { id: 47, name: '40万円', value: 40, created_at: '', updated_at: '' },
    { id: 48, name: '50万円', value: 50, created_at: '', updated_at: '' },
    { id: 49, name: '100万円', value: 100, created_at: '', updated_at: '' },
  ],
  // 駅徒歩
  walkFromStation: [
    { id: 1, name: '1分', value: 1, created_at: '', updated_at: '' },
    { id: 2, name: '5分', value: 5, created_at: '', updated_at: '' },
    { id: 3, name: '7分', value: 7, created_at: '', updated_at: '' },
    { id: 4, name: '10分', value: 10, created_at: '', updated_at: '' },
    { id: 5, name: '15分', value: 15, created_at: '', updated_at: '' },
    { id: 6, name: '20分', value: 20, created_at: '', updated_at: '' },
  ],
  // 占有面積
  occupiedArea: [
    { id: 1, name: '20m2', value: 20, created_at: '', updated_at: '' },
    { id: 2, name: '25m2', value: 25, created_at: '', updated_at: '' },
    { id: 3, name: '30m2', value: 30, created_at: '', updated_at: '' },
    { id: 4, name: '35m2', value: 35, created_at: '', updated_at: '' },
    { id: 5, name: '40m2', value: 40, created_at: '', updated_at: '' },
    { id: 6, name: '45m2', value: 45, created_at: '', updated_at: '' },
    { id: 7, name: '50m2', value: 50, created_at: '', updated_at: '' },
    { id: 8, name: '55m2', value: 55, created_at: '', updated_at: '' },
    { id: 9, name: '60m2', value: 60, created_at: '', updated_at: '' },
    { id: 10, name: '65m2', value: 100, created_at: '', updated_at: '' },
    { id: 11, name: '70m2', value: 20, created_at: '', updated_at: '' },
    { id: 12, name: '80m2', value: 80, created_at: '', updated_at: '' },
    { id: 13, name: '90m2', value: 90, created_at: '', updated_at: '' },
    { id: 14, name: '100m2', value: 100, created_at: '', updated_at: '' },
  ],
  // 築年数
  buildingAge: [
    { id: 1, name: '新築', value: 0, created_at: '', updated_at: '' },
    { id: 2, name: '1年以内', value: 1, created_at: '', updated_at: '' },
    { id: 3, name: '3年以内', value: 3, created_at: '', updated_at: '' },
    { id: 4, name: '5年以内', value: 5, created_at: '', updated_at: '' },
    { id: 5, name: '7年以内', value: 7, created_at: '', updated_at: '' },
    { id: 6, name: '10年以内', value: 10, created_at: '', updated_at: '' },
    { id: 7, name: '15年以内', value: 15, created_at: '', updated_at: '' },
    { id: 8, name: '20年以内', value: 20, created_at: '', updated_at: '' },
    { id: 9, name: '25年以内', value: 25, created_at: '', updated_at: '' },
    { id: 10, name: '30年以内', value: 30, created_at: '', updated_at: '' },
  ],
}

export const allDataState = atom({
  key: 'allData',
  default: [
    {
      user: {
        id: 1,
        name: 'ナツメグ　太郎',
        email: 'nutmeg-taro@email.com',
        password: 'hogehoge',
        companyInfo: null,
        created_at: '',
        updated_at: '',

        personalInfo: {
          id: 1,
          familyName: 'ナツメグ',
          firstName: '太郎',
          familyNameKana: 'ナツメグ',
          firstNameKana: 'タロウ',
          birthDay: '20000101',
          phoneNumber: '00011112222',
          userId: 1,
          created_at: '',
          updated_at: '',
        },

        watchList: {
          id: 1,
          description: '',
          isPurchase: false,
          userId: 1,
          detailId: 1,
          created_at: '',
          updated_at: '',
        },

        request: {
          id: 1,
          description: '',
          isPurchase: false,
          userId: 1,
          detailId: 1,
          created_at: '',
          updated_at: '',

          matching: {
            id: 1,
            requestId: 1,
            buyerUserId: 1,
            sellerUserId: 2,
            status: 'wait matching',
            sellerMessage:
              '新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。',
            created_at: '',
            updated_at: '',
          },

          detail: {
            id: 1,
            created_at: '',
            updated_at: '',

            areas: [
              {
                id: 1,
                detailId: 1,
                postCode: '9400000',
                prefecture: '新潟県',
                city: '長岡市',
                addressNumber: '',
                buildingName: '',
              },
              // {
              //   id: 2,
              //   detailId: 1,
              //   postCode: '9401000',
              //   prefecture: '新潟県',
              //   city: '長岡市',
              //   addressNumber: '',
              //   buildingName: '',
              // },
              // {
              //   id: 3,
              //   detailId: 1,
              //   postCode: '9510000',
              //   prefecture: '新潟県',
              //   city: '新潟市',
              //   addressNumber: '',
              //   buildingName: '',
              // },
            ],

            detailLimits: [
              {
                id: 1,
                classificationName: '部屋',
                limits: [
                  // {
                  //   id: 1,
                  //   name: '占有面積',
                  //   minValue: seedData.occupiedArea[0],
                  //   maxValue: seedData.occupiedArea[2],
                  // },
                  {
                    id: 2,
                    name: '賃料',
                    minValue: seedData.price[0],
                    maxValue: seedData.price[4],
                  },
                ],
              },
            ],

            detailValues: [
              // {
              //   id: 1,
              //   classificationName: '立地',
              //   values: [
              //     {
              //       id: 1,
              //       name: '駅徒歩',
              //       value: seedData.walkFromStation[2],
              //     },
              //   ],
              // },
              // {
              //   id: 2,
              //   classificationName: '建物',
              //   values: [
              //     {
              //       id: 1,
              //       name: '築年数',
              //       value: seedData.buildingAge[2],
              //     },
              //   ],
              // },
            ],

            detailTags: [
              {
                id: 1,
                classificationName: '賃料',
                tags: [seedData.rentalPrice[2]],
              },
              {
                id: 2,
                classificationName: '間取り',
                tags: [
                  seedData.roomLayout[1],
                  seedData.roomLayout[2],
                ],
              },
            ],
          },
        },
      },
    },

    // {
    //   user: {
    //     id: 2,
    //     name: 'ナツメグ　二郎',
    //     email: 'nutmeg-jiro@email.com',
    //     password: 'hogehoge',
    //     personalInfo: null,
    //     created_at: '',
    //     updated_at: '',

    //     companyInfo: {
    //       id: 1,
    //       name: '株式会社ナツメグ',
    //       phoneNumber: '00011112222',
    //       postCode: '9400001',
    //       prefecture: '新潟県',
    //       city: '長岡市',
    //       addressNumber: '1-1-1',
    //       buildingName: '長岡ビル',
    //       website: 'https://www.nutmeg.com',
    //       userId: 2,
    //     },

    //     watchList: {
    //       id: 2,
    //       description: '',
    //       isPurchase: false,
    //       userId: 2,
    //       detailId: 2,
    //       created_at: '',
    //       updated_at: '',
    //     },

    //     request: {
    //       id: 2,
    //       description: '',
    //       isPurchase: true,
    //       userId: 2,
    //       detailId: 2,
    //       created_at: '',
    //       updated_at: '',

    //       matching: {
    //         id: 2,
    //         requestId: 2,
    //         buyerUserId: 2,
    //         sellerUserId: 1,
    //         status: 'matching',
    //         sellerMessage:
    //           '新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。',
    //         created_at: '',
    //         updated_at: '',
    //       },

    //       detail: {
    //         id: 2,
    //         created_at: '',
    //         updated_at: '',

    //         areas: [
    //           {
    //             id: 4,
    //             detailId: 2,
    //             postCode: '9400001',
    //             prefecture: '新潟県',
    //             city: '長岡市',
    //             addressNumber: '1-1-1',
    //             buildingName: '長岡ビル',
    //           },
    //           {
    //             id: 5,
    //             detailId: 2,
    //             postCode: '9401000',
    //             prefecture: '新潟県',
    //             city: '長岡市',
    //             addressNumber: '',
    //             buildingName: '',
    //           },
    //         ],

    //         detailLimits: [
    //           {
    //             id: 2,
    //             classificationName: '部屋',
    //             limits: [
    //               {
    //                 id: 3,
    //                 name: '占有面積',
    //                 minValue: seedData.occupiedArea[1],
    //                 maxValue: seedData.occupiedArea[3],
    //               },
    //               {
    //                 id: 4,
    //                 name: '賃料',
    //                 minValue: seedData.price[1],
    //                 maxValue: seedData.price[3],
    //               },
    //             ],
    //           },
    //         ],

    //         detailValues: [
    //           {
    //             id: 3,
    //             classificationName: '立地',
    //             values: [
    //               {
    //                 id: 3,
    //                 name: '駅徒歩',
    //                 value: seedData.walkFromStation[3],
    //               },
    //             ],
    //           },
    //           {
    //             id: 4,
    //             classificationName: '建物',
    //             values: [
    //               {
    //                 id: 4,
    //                 name: '築年数',
    //                 value: seedData.buildingAge[4],
    //               },
    //             ],
    //           },
    //         ],

    //         detailTags: [
    //           {
    //             id: 2,
    //             classificationName: '賃料',
    //             tags: [seedData.rentalPrice[0]],
    //           },
    //           {
    //             classificationName: '間取り',
    //             tags: [seedData.roomLayout[0], seedData.roomLayout[1]],
    //           },
    //         ],
    //       },
    //     },
    //   },
    // },
    // {
    //   user: {
    //     id: 3,
    //     name: 'ナツメグ　三郎',
    //     email: 'nutmeg-saburo@email.com',
    //     password: 'hogehoge',
    //     companyInfo: null,
    //     created_at: '',
    //     updated_at: '',

    //     personalInfo: {
    //       id: 3,
    //       familyName: 'ナツメグ',
    //       firstName: '三郎',
    //       familyNameKana: 'ナツメグ',
    //       firstNameKana: 'サブロウ',
    //       birthDay: '20000101',
    //       phoneNumber: '00011112222',
    //       userId: 3,
    //       created_at: '',
    //       updated_at: '',
    //     },

    //     watchList: {
    //       id: 3,
    //       description: '',
    //       isPurchase: false,
    //       userId: 3,
    //       detailId: 3,
    //       created_at: '',
    //       updated_at: '',
    //     },

    //     request: {
    //       id: 3,
    //       description: '',
    //       isPurchase: false,
    //       userId: 3,
    //       detailId: 3,
    //       created_at: '',
    //       updated_at: '',

    //       matching: {
    //         id: 3,
    //         requestId: 3,
    //         buyerUserId: 3,
    //         sellerUserId: null,
    //         status: 'wait matching',
    //         sellerMessage:
    //           '新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。',
    //         created_at: '',
    //         updated_at: '',
    //       },

    //       detail: {
    //         id: 3,
    //         created_at: '',
    //         updated_at: '',

    //         areas: [
    //           {
    //             id: 6,
    //             detailId: 3,
    //             postCode: '9400001',
    //             prefecture: '新潟県',
    //             city: '長岡市',
    //             addressNumber: '1-1-1',
    //             buildingName: '長岡ビル',
    //           },
    //           {
    //             id: 7,
    //             detailId: 3,
    //             postCode: '9401000',
    //             prefecture: '新潟県',
    //             city: '長岡市',
    //             addressNumber: '',
    //             buildingName: '',
    //           },
    //           {
    //             id: 8,
    //             detailId: 3,
    //             postCode: '9510000',
    //             prefecture: '新潟県',
    //             city: '新潟市',
    //             addressNumber: '',
    //             buildingName: '',
    //           },
    //         ],

    //         detailLimits: [
    //           {
    //             id: 3,
    //             classificationName: '部屋',
    //             limits: [
    //               {
    //                 id: 5,
    //                 name: '占有面積',
    //                 minValue: seedData.occupiedArea[2],
    //                 maxValue: seedData.occupiedArea[5],
    //               },
    //               {
    //                 id: 6,
    //                 name: '賃料',
    //                 minValue: seedData.price[2],
    //                 maxValue: seedData.price[5],
    //               },
    //             ],
    //           },
    //         ],

    //         detailValues: [
    //           {
    //             id: 5,
    //             classificationName: '立地',
    //             values: [
    //               {
    //                 id: 5,
    //                 name: '駅徒歩',
    //                 value: seedData.walkFromStation[5],
    //               },
    //             ],
    //           },
    //           {
    //             id: 6,
    //             classificationName: '建物',
    //             values: [
    //               {
    //                 id: 6,
    //                 name: '築年数',
    //                 value: seedData.buildingAge[6],
    //               },
    //             ],
    //           },
    //         ],

    //         detailTags: [
    //           {
    //             id: 5,
    //             classificationName: '賃料',
    //             tags: [seedData.rentalPrice[1]],
    //           },
    //           {
    //             id: 6,
    //             classificationName: '間取り',
    //             tags: [seedData.roomLayout[1], seedData.roomLayout[2]],
    //           },
    //         ],
    //       },
    //     },
    //   },
    // },

    // {
    //   user: {
    //     id: 4,
    //     name: 'ナツメグ　四郎',
    //     email: 'nutmeg-shiro@email.com',
    //     password: 'hogehoge',
    //     personalInfo: null,
    //     created_at: '',
    //     updated_at: '',

    //     companyInfo: {
    //       id: 1,
    //       name: '株式会社ナツメグ',
    //       phoneNumber: '00011112222',
    //       postCode: '9400001',
    //       prefecture: '新潟県',
    //       city: '長岡市',
    //       addressNumber: '1-1-1',
    //       buildingName: '長岡ビル',
    //       website: 'https://www.nutmeg.com',
    //       userId: 4,
    //     },

    //     watchList: {
    //       id: 4,
    //       description: '',
    //       isPurchase: false,
    //       userId: 4,
    //       detailId: 4,
    //       created_at: '',
    //       updated_at: '',
    //     },

    //     request: {
    //       id: 4,
    //       description: '',
    //       isPurchase: true,
    //       userId: 4,
    //       detailId: 4,
    //       created_at: '',
    //       updated_at: '',

    //       matching: {
    //         id: 4,
    //         requestId: 4,
    //         buyerUserId: null,
    //         sellerUserId: 4,
    //         status: 'wait matching',
    //         sellerMessage:
    //           '新潟県長岡市で1Kの物件をお探しのあなたにオススメな物件があるのでご紹介します。ご興味があればご連絡ください。',
    //         created_at: '',
    //         updated_at: '',
    //       },

    //       detail: {
    //         id: 4,
    //         created_at: '',
    //         updated_at: '',

    //         areas: [
    //           {
    //             id: 9,
    //             detailId: 4,
    //             postCode: '9400001',
    //             prefecture: '新潟県',
    //             city: '長岡市',
    //             addressNumber: '1-1-1',
    //             buildingName: '長岡ビル',
    //           },
    //           {
    //             id: 10,
    //             detailId: 4,
    //             postCode: '9401000',
    //             prefecture: '新潟県',
    //             city: '長岡市',
    //             addressNumber: '',
    //             buildingName: '',
    //           },
    //         ],

    //         detailLimits: [
    //           {
    //             id: 4,
    //             classificationName: '部屋',
    //             limits: [
    //               {
    //                 id: 7,
    //                 name: '占有面積',
    //                 minValue: seedData.occupiedArea[4],
    //                 maxValue: seedData.occupiedArea[7],
    //               },
    //               {
    //                 id: 8,
    //                 name: '賃料',
    //                 minValue: seedData.price[4],
    //                 maxValue: seedData.price[7],
    //               },
    //             ],
    //           },
    //         ],

    //         detailValues: [
    //           {
    //             id: 7,
    //             classificationName: '立地',
    //             values: [
    //               {
    //                 id: 7,
    //                 name: '駅徒歩',
    //                 value: seedData.walkFromStation[7],
    //               },
    //             ],
    //           },
    //           {
    //             id: 8,
    //             classificationName: '建物',
    //             values: [
    //               {
    //                 id: 8,
    //                 name: '築年数',
    //                 value: seedData.buildingAge[6],
    //               },
    //             ],
    //           },
    //         ],

    //         detailTags: [
    //           {
    //             id: 7,
    //             classificationName: '賃料',
    //             tags: [seedData.rentalPrice[4]],
    //           },
    //           {
    //             id: 8,
    //             classificationName: '間取り',
    //             tags: [seedData.roomLayout[4], seedData.roomLayout[5]],
    //           },
    //         ],
    //       },
    //     },
    //   },
    // },
  ],
  effects_UNSTABLE: [persistAtom],
})
