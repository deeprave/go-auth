import {
  Avatar,
  Box,
  Button,
  Checkbox,
  FormControlLabel,
  Grid,
  Link,
  Paper,
  TextField,
  Typography,
} from "@suid/material"
import { Email, LockOutlined } from "@suid/icons-material"
import { Link as RouterLink } from "@solidjs/router"

import Copyright from "../components/Copyright.js"
import Apple from "../icons/Apple"
import Google from "../icons/Google"

//import { useAppContext } from "../utils/AppContext"

export default function SignIn() {
  // const { setJwtToken } = useAppContext()

  const handleSubmit = (event: SubmitEvent) => {
    event.preventDefault()
    const data = new FormData(event.currentTarget as HTMLFormElement)
    console.log({
      username: data.get('username'),
      password: data.get('password'),
    })
  }

  return (
    <Grid container component="main" sx={{ height: '100vh' }}>
      <Grid item xs={false} sm={4} md={7} sx={{
        backgroundImage: 'url(src/assets/Network-Safety-Concept.jpg)',
        backgroundRepeat: 'no-repeat',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
      }}/>
      <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
        <Box sx={{ my: 2, mx: 4, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
          <Avatar sx={{ m: 1, bgcolor: 'primary.main' }}>
            <LockOutlined/>
          </Avatar>
          <Typography component="h1" variant="h4">
            Sign In
          </Typography>
          <Button fullWidth variant="contained" sx={{ mt: 1, backgroundColor: "red" }}>
            <Google/>&nbsp;
            Sign In using Google
          </Button>
          <Button name="apple" fullWidth variant="contained" sx={{ mt: 1, backgroundColor: "green" }}>
            <Apple/>&nbsp;
            Sign In using Apple
          </Button>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 1 }}>
            <Button fullWidth variant="contained" sx={{ mt: 1, mb: 1 }} disabled>
              <Email/>&nbsp;
              Sign In using email
            </Button>
            <TextField name="username" margin="normal" required fullWidth
                       label="Username or Email Address" autoComplete="username" autoFocus/>
            <TextField name="password" type="password" margin="normal" required fullWidth
                       label="Password" autoComplete="current-password"/>
            <FormControlLabel control={<Checkbox value="remember" color="primary"/>} label="Remember me"/>
            <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
              Sign In
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                Don't have an account?
                <RouterLink href="/signup">
                  <Button variant="text">Sign Up</Button>
                </RouterLink>
              </Grid>
            </Grid>
            <Copyright sx={{ mt: 5, flexGrow: 1 }}/>
          </Box>
        </Box>
      </Grid>
    </Grid>
  )
}
