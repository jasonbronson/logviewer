module.exports = {
  root: false,
  env: {
    node: true
  },
  extends: ["plugin:vue/base", 'eslint:recommended'],
  parserOptions: {
    parser: 'babel-eslint'
  },
  rules: {
    'no-console': 'off',
    'no-debugger':
      process.env.ENVIRONMENT_NAME === 'production' ? 'error' : 'off',
    'no-prototype-builtins': ['off'],
    'no-empty': ['off'],
    'no-unused-vars': ['warn']
  }
}
