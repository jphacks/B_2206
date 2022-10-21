import { atom } from 'recoil'
import { recoilPersist } from 'recoil-persist'

const { persistAtom } = recoilPersist()

export const conditionState = atom({
  key: 'condition',
  default: {
    id: 1,
    mode: '',
    prefectureName: '',
    prefectureId: '',
    cityNames: [],
    conditions: {},
    isPostCode: false,
  },
  effects_UNSTABLE: [persistAtom],
})
