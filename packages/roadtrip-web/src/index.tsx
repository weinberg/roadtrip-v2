import React, { FC } from 'react';
import { render } from 'react-dom';
import { ApolloClient, ApolloProvider, gql, InMemoryCache, useQuery } from '@apollo/client';

console.log(`uri: ${process.env.API_URI}`);

const client = new ApolloClient({
  uri: 'http://localhost:8080',
  cache: new InMemoryCache(),
  headers: {
    Authorization: '49ba3caa-d0da-4e81-9ee0-47ce29d05e69',
  },
});

const GET_CHARACTER = gql`
  query CurrentCharacter {
    currentCharacter {
      id
      name
      car {
        id
        mph
        name
        odometer
        tripometer
        route {
          id
          name
        }
        location {
          index
          miles
          routeId
        }
      }
    }
  }
`;

const GET_CAR_STATUS = gql`
  query Car($id: ID!) {
    car(id: $id) {
      id
      mph
      odometer
    }
  }
`;

const Character: FC = () => {
  const { loading, error, data: { currentCharacter } = {} } = useQuery(GET_CHARACTER);
  const { data: carData } = useQuery(GET_CAR_STATUS, {
    variables: { id: currentCharacter?.car?.id },
    skip: !currentCharacter?.car?.id,
    pollInterval: 1000,
  });

  return (
    <>
      <div>loading: {loading?.toString()}</div>
      <div>error: {error?.toString()}</div>
      <div>currentCharacter: {JSON.stringify(currentCharacter, null, 2)}</div>
      <div>carData: {JSON.stringify(carData, null, 2)}</div>
    </>
  );
};

function App() {
  return (
    <>
      <div>
        <h2>My first Apollo app 🚀</h2>
        <Character />
      </div>
    </>
  );
}

render(
  <ApolloProvider client={client}>
    <App />
  </ApolloProvider>,
  document.getElementById('root')
);
