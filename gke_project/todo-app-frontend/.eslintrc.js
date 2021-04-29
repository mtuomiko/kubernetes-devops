module.exports = {
  'env': {
    'es6': true,
    'browser': true,
    'node': true,
    'commonjs': true,
  },
  // 'globals': {
  //   'BACKEND_URL': true,
  // },
  'extends': [
    'eslint:recommended',
    'plugin:react/recommended',
    'plugin:@typescript-eslint/recommended',
  ],
  'parser': '@typescript-eslint/parser',
  // 'parserOptions': {
  //   'ecmaFeatures': {
  //     'jsx': true,
  //   },
  //   'ecmaVersion': 2018,
  //   'sourceType': 'module',
  // },
  'plugins': [
    'react',
    '@typescript-eslint',
  ],
  'rules': {
    'react/prop-types': 0,
    '@typescript-eslint/explicit-module-boundary-types': 0,
  },
  'settings': {
    'react': {
      'version': 'detect',
    }
  }
}