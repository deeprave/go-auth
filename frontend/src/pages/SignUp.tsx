import { Link as RouterLink } from '@solidjs/router'
import {
  Avatar,
  Button,
  TextField,
  Grid,
  Box,
  Typography,
  Paper
} from '@suid/material'
import {
  Email,
  LockOutlined
} from "@suid/icons-material"
import Google from '../icons/Google'
import Apple from '../icons/Apple'

import Copyright from '../components/Copyright'

export default function SignUp() {
  const handleSubmit = (event: SubmitEvent) => {
    event.preventDefault()
    const data = new FormData(event.currentTarget as HTMLFormElement)
    console.log({
      username: data.get('username'),
      given: data.get('given'),
      surname: data.get('surname'),
      email: data.get('email'),
      password: data.get('password'),
      confirmPassword: data.get('confirm-password'),
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
            Sign Up
          </Typography>
          <Button fullWidth variant="contained" sx={{ mt: 1, backgroundColor: "red" }}>
            <Google/>&nbsp;
            Sign Up using Google
          </Button>
          <Button name="apple" fullWidth variant="contained" sx={{ mt: 1, backgroundColor: "green" }}>
            <Apple/>&nbsp;
            Sign Up using Apple
          </Button>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 1 }}>
            <Button fullWidth variant="contained" sx={{ mt: 1, mb: 1, backgroundColor: "blue" }} disabled>
              <Email/>&nbsp;
              Sign Up using this form
            </Button>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <TextField name="username" required fullWidth autoFocus
                           label="Username" autoComplete="username"/>
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField name="given" required fullWidth
                           label="Given Name" autoComplete="first-name"/>
              </Grid>
              <Grid item xs={12} md={6}>
                <TextField name="surname" required fullWidth
                           label="Surname" autoComplete="last-name"/>
              </Grid>
              <Grid item xs={12}>
                <TextField name="email" required fullWidth
                           label="Email Address" autoComplete="email"/>
              </Grid>
              <Grid item xs={12}>
                <TextField name="password" required fullWidth type="password"
                           label="Password" autoComplete="new-password"/>
              </Grid>
              <Grid item xs={12}>
                <TextField name="confirm-password" required fullWidth type="password"
                           label="Confirm password" autoComplete="confirm-password"/>
              </Grid>
            </Grid>
            <Button type="submit" fullWidth variant="contained" sx={{ mt: 1, mb: 1 }}>
              Sign Up
            </Button>
            <Grid container justifyContent="flex-end">
              <Grid item>
                Already have an account?
                <RouterLink href="/signin">
                  <Button variant="text">Sign In</Button>
                </RouterLink>
              </Grid>
            </Grid>
            <Copyright sx={{ mt: 5 }}/>
          </Box>
        </Box>
      </Grid>
    </Grid>
  )
}
