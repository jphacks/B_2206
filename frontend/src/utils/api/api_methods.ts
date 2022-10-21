import axios from 'axios'

export const get = async (url: string) => {
  axios
    .get(url)
    .then((results) => {
      // console.log(results.data.data);
      const data: any = results.data.data
      return data
    })
    .catch((error) => {
      console.log(error)
    })
}

export const getWithSet = async (url: string, set: Function) => {
  axios
    .get(url)
    .then((results) => {
      // console.log(results.data.data);
      const data: any = results.data.data
      set(data)
    })
    .catch((error) => {
      console.log(error)
    })
}
