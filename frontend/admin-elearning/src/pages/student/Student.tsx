import useUserStore from '@/store/user'
import React, { useEffect } from 'react'

export default function Student() {
  const {setUsers , users} = useUserStore();

  useEffect(() => {
    setUsers();
  },[])

  console.log("users", users);

  return (
    <React.Fragment>
      <h1>Student</h1>
    </React.Fragment>
  )
}
