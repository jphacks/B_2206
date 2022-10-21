import axios from 'axios'
import { setDefaultResultOrder } from 'dns'

export const get = async (url: string) => {
  axios
    .get(url)
    .then((results) => {
      const data: any = results.data.data
      return data
    })
    .catch((error) => {
    })
}

export const getWithSet = async (url: string, set: Function) => {
  axios
    .get(url)
    .then((results) => {
      const data: any = results.data.data
      set(data)
    })
    .catch((error) => {
      set(error)
    })
}

export const getPostWithSet = async (url: string, set: Function) => {
  axios
    .get(url)
    .then((results) => {
      const data: any = results.data
      set(data)
    })
    .catch((error) => {
      set(error)
    })
}
