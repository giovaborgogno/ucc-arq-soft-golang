import { useState, useEffect } from 'react';
import Link from 'next/link';
const axios = require('axios')
import { useRouter } from "next/router";


const Users = ({ users }) => {
  const router = useRouter();
  const handleDelete = async (id) => {
    await axios.delete(`http://localhost:8080/users/${id}`);
    router.push("/users");
  };

  return (
    <div>
      <h1>Users</h1>
      <Link href={'/create'}> <button>Create New User</button></Link>

      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Email</th>
            <th></th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {users.map(user => (
            <tr key={user.id}>
              <td>{user.nombre}</td>
              <td>{user.email}</td>
              <td>
                <Link href={`/users/${user.id}`}>
                  Edit
                </Link>
              </td>
              <td>
                <button onClick={() => handleDelete(user.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Users;

export async function getServerSideProps(context) {
  const response = await axios.get('http://localhost:8080/users');
  const users = await response.data;
  return {
    props: { users }, // will be passed to the page component as props
  }
}