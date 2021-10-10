/**
 *  Print prisma database query with variables injected so you can run the sql manually
 *  requires 'query' to be in prisma client log configuration: `log: ['query'],`
 *  After prisma client is initialised do:
 * client.$on('query', logSql);
 */

// print query with params inserted
// eslint-disable-next-line @typescript-eslint/explicit-module-boundary-types
const logSql = (e): void => {
  let query = e.query;
  let paramString: string = e.params.replace(/[[\]]/g, '');
  paramString = paramString.replace(/"/g, "'");
  const params: string[] = paramString.split(',');

  for (let i = 0; i < e.params.length; i++) {
    let value = params[i];
    if (typeof value === 'string') {
      value = `"${value}"`;
    }
    query = query.replace(`$${i + 1}`, params[i]);
  }
  query = query.replace(/"public"\./g, '');

  //eslint-disable-next-line no-console
  console.log(query);
};

export { logSql };
