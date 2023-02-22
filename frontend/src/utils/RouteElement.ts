import { JSX } from "solid-js";

export type RouteElement = {
  icon?: JSX.Element
  label: string
  href?: string
}

export type RouteElements = RouteElement[]
