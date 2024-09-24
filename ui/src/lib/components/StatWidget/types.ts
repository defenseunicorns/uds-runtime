import { type CarbonIcon } from 'carbon-icons-svelte'

export type BarSizeType = 'sm' | 'md' | 'lg' | 'xl'
export type UnitType = 'Cores' | 'GB'
export type TailwindSizeType = 16 | 20 | 24 | 28 | 32 | 36 | 40 | 44 | 48

export type VariantType = {
  title: string
}

export type WithRightIconType = {
  subtitle: string
  icon: typeof CarbonIcon
  link: string
}

export type ProgressBarType = {
  capacity: number
  progress: number
  unit: UnitType
  value: string | number
}

export type StatType = VariantType & (WithRightIconType | ProgressBarType)
