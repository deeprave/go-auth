import { Route, Routes } from "@solidjs/router"
import {
  Box,
  CssBaseline,
  IconButton,
  ThemeProvider
} from "@suid/material"
import {
  Delete,
  Edit,
  LockClock,
  LockOpen,
  Logout,
  NavigateBefore,
  NavigateNext,
  Person as User,
  Group
} from '@suid/icons-material'
import { For, Show } from "solid-js";

import SignIn from "./pages/SignIn"
import SignUp from "./pages/SignUp"

import { AppLayout } from "./components/AppLayout"

import { useAppContext } from "./utils/AppContext";
import { RouteElements } from "./utils/RouteElement";
import { Tools } from "./utils/ToolbarItems";


export default function App() {
  const { theme } = useAppContext()

  const routes = (): RouteElements => ([
    {
      label: "Users",
      href: "/users",
      icon: <User/>
    },
    {
      label: "Groups",
      href: "/groups",
      icon: <Group/>
    },
    {
      label: "Sign In",
      href: "/signin",
      icon: <LockOpen/>,
    },
    {
      label: "Sign Up",
      href: "/signup",
      icon: <LockClock/>,
    }
  ])
  const toolbar = (): Tools => ([
    {
      icon: <Edit/>,
      tooltip: "edit",
      disabled: true,
      func: () => {}
    },
    {
      icon: <Delete/>,
      tooltip: "delete",
      disabled: true,
      func: () => {}
    },
    {
      tooltip: "",
      disabled: true,
      func: () => {}
    },
    {
      icon: <NavigateBefore/>,
      tooltip: "back",
      disabled: true,
      func: () => {}
    },
    {
      icon: <NavigateNext/>,
      tooltip: "delete",
      disabled: true,
      func: () => {}
    },
    {
      icon: <Logout/>,
      tooltip: "logout",
      disabled: false,
      func: () => {}
    }
  ])

  const userButtons = () => (
    <Box justifyContent="space-between">
      <For each={toolbar()}>
        {(tool) => {
        return (
          <Show when={tool.icon} fallback=" " keyed={false}>
            <IconButton color="inherit" disabled={tool.disabled}>{tool.icon}</IconButton>
          </Show>
        )
      }
      }</For>
    </Box>
  )
  return (
    <>
      <CssBaseline/>
      <ThemeProvider theme={theme}>
        <Routes>
          <Route path="/signin" element={<SignIn/>}/>
          <Route path="/signup" element={<SignUp/>}/>
          <Route path="/" element={<AppLayout
            title="Authentication Manager"
            buttons={userButtons()}
            routes={routes()}
          />}>
          </Route>
        </Routes>
      </ThemeProvider>
    </>
  )
}
