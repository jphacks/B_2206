export interface User {
  id: number
  name: string
  email: string
  password: string
  personalInfoId: number
  companyInfoId: number
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
  phoneNumber: number
  userId: number
  created_at: string
  updated_at: string
}

export interface CompanyInfo {
  id: number
  name: string
  phoneNumber: number
  postCode: number
  prefecture: string
  city: string
  addressNumber: string
  buildingName: string
  website: string
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

export interface Detail {
  id: number
  maxPrice: number
  minPrice: number
  maxBuildingAge: number
  minBuildingAge: number
  maxLandArea: number
  minLandArea: number
  maxBuildingArea: number
  minBuildingArea: number
  created_at: string
  updated_at: string
}

export interface DetailsAddress {
  id: number
  detailId: number
  addressId: number
  created_at: string
  updated_at: string
}

export interface Address {
  id: number
  postCode: number
  prefecture: string
  city: string
  addressNumber: string
  buildingName: string
  created_at: string
  updated_at: string
}

export interface DetailsCategory {
  id: number
  detailId: number
  categoryId: number
  created_at: string
  updated_at: string
}

export interface Category {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsClosestStation {
  id: number
  detailId: number
  closestStationId: number
  created_at: string
  updated_at: string
}

export interface ClosestStation {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsLayout {
  id: number
  detailId: number
  layoutId: number
  created_at: string
  updated_at: string
}

export interface Layout {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsPerformance {
  id: number
  detailId: number
  performanceId: number
  created_at: string
  updated_at: string
}

export interface Performance {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsReformRenovation {
  id: number
  detailId: number
  reformRenovationId: number
  created_at: string
  updated_at: string
}

export interface ReformRenovation {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsDialy {
  id: number
  detailId: number
  dialyId: number
  created_at: string
  updated_at: string
}

export interface Dialy {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsBathroom {
  id: number
  detailId: number
  bathroomId: number
  created_at: string
  updated_at: string
}

export interface Bathroom {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsBalcony {
  id: number
  detailId: number
  balconyId: number
  created_at: string
  updated_at: string
}

export interface Balcony {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsParking {
  id: number
  detailId: number
  layoutId: number
  created_at: string
  updated_at: string
}

export interface Parking {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsLayout {
  id: number
  detailId: number
  layoutId: number
  created_at: string
  updated_at: string
}

export interface Layout {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsLayout {
  id: number
  detailId: number
  layoutId: number
  created_at: string
  updated_at: string
}

export interface Layout {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface DetailsLayout {
  id: number
  detailId: number
  layoutId: number
  created_at: string
  updated_at: string
}

export interface Layout {
  id: number
  name: string
  created_at: string
  updated_at: string
}

export interface request {
  id: number
  created_at: string
  updated_at: string
}
