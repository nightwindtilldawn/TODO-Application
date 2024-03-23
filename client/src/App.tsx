import { Box } from '@mantine/core'
import useSWR from 'swr';
import './App.css'

export const ENDPOINT = 'http://localhost:4000'

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: boolean;
}

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());
 

function App() {

  const { data, mutate } = useSWR('api/todos', fetcher);

  return <Box>(JSON.stringify(data))</Box>;
  
}

export default App
