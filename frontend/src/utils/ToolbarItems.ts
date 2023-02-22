import { JSX } from "solid-js";

export type Tool = {
  icon?: JSX.Element
  tooltip: string
  disabled: boolean
  func: () => void
}

export type Tools = Tool[]
