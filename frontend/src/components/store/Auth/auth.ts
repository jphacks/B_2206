import { atom } from 'recoil'
import { recoilPersist } from 'recoil-persist'

const { persistAtom } = recoilPersist()

export const userState = atom({
  key: 'user',
  default: {
    id: 1,
    name: 'gidai yuuki',
    email: 'gidai@email',
  },
  effects_UNSTABLE: [persistAtom],
})
