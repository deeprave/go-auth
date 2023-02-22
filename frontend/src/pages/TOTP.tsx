import {
  Avatar,
  Box,
  Button,
  Grid,
  Paper,
  TextField,
  Typography
} from "@suid/material";
import { LockOutlined } from "@suid/icons-material";
import Copyright from "../components/Copyright";

export default function TOTP() {

  const handleSubmit = (event: SubmitEvent) => {
    event.preventDefault()
    const form = event.currentTarget as HTMLFormElement
    const data = new FormData(form)
    console.log({
      email: data.get('totp'),
    })
  }

  return (
    <Grid container component="main" sx={{ height: '100vh' }}>
      <Grid item xs={false} sm={4} md={7} sx={{
        backgroundImage: 'url(src/assets/Network-Safety-Concept.jpg)',
        backgroundRepeat: 'no-repeat',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
      }}
      />
      <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
        <Box sx={{ my: 2, mx: 4, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
          <Avatar sx={{ m: 1, bgcolor: 'primary.main' }}>
            <LockOutlined/>
          </Avatar>
          <Typography component="h1" variant="h4">
            Enter Authenticator Code
          </Typography>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 1 }}>
            <TextField name="token" type="number" margin="normal" fullWidth
                       label="Authenticator Token" autoComplete="token" inputMode="numeric"/>
            <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
              Confirm
            </Button>
            <Copyright sx={{ mt: 5 }}/>
          </Box>
        </Box>
      </Grid>
    </Grid>
  )
}
