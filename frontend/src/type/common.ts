export interface User {
  id: number
  name: string
  email: string
  password: string
  created_at: string
  updated_at: string
}

export interface PersonalInfo {
  id: number
  familyName: string
  firstName: string
  familyNameKana: string
  firstNameKana: string
  birthDay: string
  phoneNumber: string
  userId: number
  created_at: string
  updated_at: string
}

export interface CompanyInfo {
  id: number
  name: string
  phoneNumber: string
  postCode: string
  prefecture: string
  city: string
  addressNumber: string
  buildingName: string
  website: string
  userId: number
  created_at: string
  updated_at: string
}

export interface Request {
  id: number
  description: string
  isPurchase: boolean
  userId: number
  detailId: number
  created_at: string
  updated_at: string
}

export interface WatchList {
  id: number
  description: string
  isPurchase: boolean
  userId: number
  detailId: number
  created_at: string
  updated_at: string
}

export interface Matching {
  id: number
  requestId: number
  buyerUserId: number
  sellerUserId: number
  status: string
  sellerMessage: string
  created_at: string
  updated_at: string
}

export interface Wish {
  id: number
  requestId: number
  userId: number
  created_at: string
  updated_at: string
}

export interface Detail {
  id: number
  areaId: number
  created_at: string
  updated_at: string
}

export interface Area {
  id: number
  postCode: number
  prefecture: string
  city: string
  addressNumber: string
  buildingName: string
  created_at: string
  updated_at: string
}

export interface Classification {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsValue {
  id: number
  detailId: number
  valueId: number
  classificationId: number
  created_at: string
  updated_at: string
}

export interface Value {
  id: number
  name: string
  value: number
  limitId: number
  created_at: string
  updated_at: string
}

export interface DetailsLimit {
  id: number
  detailId: number
  limitId: number
  classificationId: number
  created_at: string
  updated_at: string
}

export interface Limit {
  id: number
  name: string
  maxValueId: number
  minValueId: number
  created_at: string
  updated_at: string
}

export interface DetailsTag {
  id: number
  detailId: number
  tagId: number
  classificationId: number
  created_at: string
  updated_at: string
}

export interface Tag {
  id: number
  name: string
  created_at: string
  updated_at: string
}
