import { createSignal, createContext, useContext, JSX } from "solid-js"
import { createTheme, Theme } from "@suid/material/styles"

export type AppContextType = {
  theme: Theme
  getJwtToken: () => string
  setJwtToken: (token: string) => boolean
  isLoggedIn: () => boolean
}

export const AppContext = createContext<AppContextType|null>(null)

export type AppContextProps = {
  children?: JSX.Element
}

export const AppContextProvider = (props: AppContextProps) => {
  const [getJwtToken, setJwtToken] = createSignal<string>("")

  function isValidToken(token: string): [boolean, number] {
    // todo: fix me, validate jwt token and extract expiry time
    token
    return [true, Date.now() + (60 * 60 * 1000)]
  }

  function appTheme(): Theme {
    return createTheme({
      palette: {
        action: {
          disabledBackground: "black",
          disabled: "lightblue"
        }
      }
    })
  }

  return (
      <AppContext.Provider value={{
        theme: appTheme(),
        getJwtToken: () => getJwtToken(),
        setJwtToken: (token: string) => {
          if (isValidToken(token)) {
            setJwtToken(token)
            return true
          }
          return false
        },
        isLoggedIn: () => getJwtToken() != "",
      }}>
        {props.children}
      </AppContext.Provider>
  )
}

export const useAppContext = (): AppContextType => {
  const appContext = useContext(AppContext)

  if (!appContext)
    throw new Error("appContext used outside of AppContextProvider block")
  return appContext
}
