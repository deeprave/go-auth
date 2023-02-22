# Auth Manager
This is a simple application used to manage users, groups and user credentials.

## Setup

```shell
$ pnpm install # or npm install or yarn install
```

## Background
This app is written in solid-js.
You can learn more on the [Solid Website](https://solidjs.com) and chat on [Discord](https://discord.com/invite/solidjs)

The current app was originally written with
- React v18.x
- React Router v6, and
- MaterialUI v5.x

It has been re-imagined and translated to
- SolidJS 1.6
- SolidJS Router v0.7
- SUID material and material-icons

Both applications used [Vite](https://vitejs.dev) as the bundler, development platform and transpiler.
The original ReactJS app also used [Vitest](https://vitest.dev) with @testing-library/react, moving tests
to SolidJS is currently in progress.

The backend used in both cases is identical, written in Golang for simplicity and performance.


### Why?

The question is: why to move this app from ReactJS to Solid?

The two libraries are most definitely NOT compatible, although both use a component model and - especially -
JSX as the primary templating language.
This makes them very similar in application structure. There are, however, some very fundamental differences,
and knowing them will help answer the question.

#### The virtual DOM

- ReactJS utilizes a technique called reconciliation, which involves maintaining a "virtual" 
representation of the user interface in memory known as the Virtual DOM, or VDOM.
Access to the VDOM is extremely fast, compared with the actual DOM where operations are slower as
they involve direct and indirect interaction with the browser/
The "virtual" DOM is then diffed against the actual DOM, and only any necessary updates are made to the actual
DOM, enough to keep them in sync and minimise costly DOM manipulation to improve overall performance.

- Instead, SolidJS utilizes a unique approach to updating the UI. Instead of using a Virtual DOM, Solid JS compiles
templates into actual DOM nodes, which are then updated with specific reactions when the application's state changes.
The need for reconciliation is removed and the DOM is updated directly, resulting in faster updates.
A state can be declared and used throughout the application, and only the code that depends on it will 
be rerun when the state value is modified. This helps minimize state change's impact on the UI, further 
improving the speed and efficiency of the application.

#### Components

- A solidJS component is called only once and the DOM is instantiated accordingly/ 

- In React, we can use useEffect function for init call, cleanup, and reaction to data changes, but SolidJS, 
we have onMount for init, onCleanup for cleanup, and createEffect to react to data changes.

#### Reactivity

- ReactJS reactive data are raw values. Every time a state value is changed, the whole component/function is re-rendered.

- In Solid, you have functions instead of raw values. That's why you need to call them to get the current state value. 
You can think of a component as a constructor function.
How is the view updated when the component is NOT called again with the new values?
Well, when you want to display the value, you need to call it, so the reactive value is a getter function.
To accomplish this, SolidJS uses signals, which are event emitters that hold a list of subscriptions.
They notify their subscribers whenever their value changes.

#### Lifecycle Functions

- In React JS, as the init and cleanup function, we can use `useEffect(()`, and have it return a function that is
run as cleanup.
The function returned from the callback is called whenever the component is unmounted.

- In Solid JS, the lifecycle functions are pretty self-explanatory:
  - `onMount(()` is called when the component is fully initialised
  - `onCleanup(()` is called when the component is unmounted

#### Effect Functions

- In React JS, to react to data changes, we can use `useEffect()` function.
The callback inside useEffect is called as many times as the state value changes.
Also, interestingly, before each render, the cleanup function is called as well.

- In Solid JS, this process is more straightforward.
`createEffect()` callback function is called as many times as the state value changes.
Solid JS doesn't require or have a dependency array.
Instead, dependencies are detected automatically whenever the signal values are used.

#### State

