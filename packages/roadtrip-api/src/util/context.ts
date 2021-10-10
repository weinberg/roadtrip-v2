import { Character } from '../resolvers/character';
import { PrismaClient } from '@prisma/client';
import { logError, logInfo } from './logging';

interface Context {
  currentCharacter: Character;
  db: PrismaClient;
}

const getContext = async (db: PrismaClient, token: string): Promise<Context> => {
  logInfo(`GetContext finding character by token: ${token}`);
  let currentCharacter: Character;
  try {
    currentCharacter = await db.character.findUnique({ where: { token } });
  } catch (e) {
    logError(e.toString());
  }
  logInfo(`currentCharacter: ` + JSON.stringify(currentCharacter));
  return {
    currentCharacter,
    db,
  };
};

export { Context, getContext };
