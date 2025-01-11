import React from "react"
import { createRoot } from "react-dom/client"
import { Provider } from "react-redux"
import { store } from "./app/store"
import axios from "axios"

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import AppRouter from "./AppRouter";
import { CssBaseline, GlobalStyles, ThemeProvider } from "@mui/material"
import theme from "./theme"


// const { VITE_APP_SERVER_URL } = import.meta.env;

const root = createRoot(document.body);

axios.defaults.baseURL = 'http://localhost:8000'
axios.defaults.withCredentials = true
// console.log('VITE_APP_SERVER_URL', VITE_APP_SERVER_URL)
// console.log('process.env', import.meta.)

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <GlobalStyles
          styles={{
            body: {
              background: 'linear-gradient(135deg, #010101 0%, #402b4c 100%)',
              minHeight: '100vh',
              margin: 0,
            },
          }}
        />
        <div style={{ position: 'relative', width: '100%', display: 'flex', justifyContent: 'center', alignItems: 'center', flexDirection: 'column' }}>
          <AppRouter />
        </div>
      </ThemeProvider>
    </Provider>
  </React.StrictMode>,
)
