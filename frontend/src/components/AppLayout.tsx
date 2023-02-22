import {
  AppBar,
  Box,
  Button,
  Drawer,
  Toolbar,
  Typography,
} from "@suid/material"
import { Menu, MenuOpen } from "@suid/icons-material"
import { Outlet } from "@solidjs/router"
import { createSignal, JSX, Show } from "solid-js"

import SideBar from './SideBar'
import classes from './AppLayout.module.scss'
import { RouteElements } from "../utils/RouteElement";


export type AppLayoutProps = {
  title: string
  buttons: JSX.Element
  routes: RouteElements
}

export function AppLayout(props: AppLayoutProps) {
  const [isOpen, setOpen] = createSignal<boolean>(false)

  const excludeEventKeys = [
    'Tab', 'Shift', 'Control', 'ArrowUp', 'ArrowDown', 'ArrowLeft', 'ArrowRight'
  ]
  const toggleDrawer = (open: boolean) => (event: KeyboardEvent | MouseEvent) => {
    if (event.type === 'keydown') {
      const keyEvent = event as KeyboardEvent
      console.log(keyEvent.key)
      if (excludeEventKeys.indexOf(keyEvent.key) != -1) {
        return;
      }
    }
    setOpen(open)
  }

  const sideBar = () => (
    <SideBar
      title={props.title}
      onClick={toggleDrawer(false)}
      onKeyDown={toggleDrawer(false)}
      routes={props.routes}
    />
  )
  return (
    <>
      <AppBar position="sticky">
        <Toolbar>
          <Box flex="row" flexGrow={1}>
            <Typography as="title" variant="h4" class={classes.AppTitle}>
              <Button color="inherit" onClick={toggleDrawer(true)}>
                <Show when={isOpen()} fallback={<Menu/>} keyed={false}><MenuOpen/></Show>
              </Button>
              {props.title}
            </Typography>
          </Box>
          <Box class={classes.inactive}>
            {props.buttons}
          </Box>
        </Toolbar>
      </AppBar>
      <Drawer
        open={isOpen()}
        onClose={toggleDrawer(false)}
      >
        {sideBar()}
      </Drawer>
      <Outlet/>
    </>
  )
}
