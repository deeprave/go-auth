import Typography from "@suid/material/Typography"
import Link from "@suid/material/Link"

export default function Copyright(props: any) {
  return (
    <Typography variant="body2" color="text.secondary" align="center" marginTop="auto" {...props}>
      {'Copyright © '}
      <Link color="inherit" href="https://uniquode.io/">
        Uniquode
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  )
}
