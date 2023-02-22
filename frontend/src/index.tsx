/* @refresh reload */
import { render } from 'solid-js/web'
import { Router } from '@solidjs/router'

import App from './App'
import { AppContextProvider } from "./utils/AppContext"

const root = document.getElementById('root')
render(() =>
    <AppContextProvider>
      <Router>
        <App/>
      </Router>
    </AppContextProvider>
, root!)
