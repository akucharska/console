import DocsContent from './DocsContent/resolvers';

export default {
  Mutation: {
    ...DocsContent.Mutation,
  },
  Query: {
    ...DocsContent.Query,
  },
};
