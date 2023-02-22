import { For, Show } from "solid-js";
import { Box, List, ListItem, ListItemButton, ListItemText, Typography } from "@suid/material";
import { RouteElements } from "../utils/RouteElement";


export type SideBarProps = {
  title?: string
  onClick?: (event: MouseEvent | KeyboardEvent) => void
  onKeyDown?: (event: MouseEvent | KeyboardEvent) => void
  routes: RouteElements
}

const SideBar = (props: SideBarProps) => (
  <Box
    sx={{
      width: 250,
      backgroundColor: "#eef",
      height: "100vh",
      padding: 1
  }}
    color="secondary"
    role="presentation"
    onClick={props.onClick}
    onKeyDown={props.onKeyDown}
  >
    <Show when={props.title} keyed={false}>
      <Box sx={{textAlign: 'center'}}>
        <Typography as="h2" variant="h6">
          {props.title}
        </Typography>
      </Box>
    </Show>
    <List disablePadding>
      <For each={props.routes}>
        {(route) =>
          <ListItem autoFocus={true} disablePadding dense>
            <ListItemButton>
              {route.icon}&nbsp;<ListItemText primary={route.label}/>
            </ListItemButton>
          </ListItem>
        }</For>
    </List>
  </Box>
)

export default SideBar
