// Sample function data matching the FunctionComponent schema
export const sampleFunctions = [
  {
    id: 'func-1',
    name: 'calculateTotal',
    visibility: 'public',
    description: 'Calculates the total sum of items in a cart',
    position: { x: 100, y: 100 },
    sourceLocation: {
      file: 'src/cart.js',
      startLine: 15,
      endLine: 28
    },
    signature: {
      name: 'calculateTotal',
      parameters: [
        {
          name: 'items',
          type: 'Array<CartItem>',
          description: 'Array of cart items',
          optional: false
        },
        {
          name: 'taxRate',
          type: 'number',
          description: 'Tax rate as decimal',
          optional: true,
          defaultValue: '0.08'
        }
      ],
      returns: [
        {
          type: 'number',
          description: 'Total price including tax'
        }
      ]
    }
  },
  {
    id: 'func-2',
    name: 'validateUser',
    visibility: 'private',
    description: 'Validates user credentials against database',
    position: { x: 450, y: 100 },
    sourceLocation: {
      file: 'src/auth.js',
      startLine: 42,
      endLine: 67
    },
    signature: {
      name: 'validateUser',
      parameters: [
        {
          name: 'username',
          type: 'string',
          description: 'User username',
          optional: false
        },
        {
          name: 'password',
          type: 'string',
          description: 'User password (hashed)',
          optional: false
        }
      ],
      returns: [
        {
          type: 'Promise<boolean>',
          description: 'True if credentials are valid'
        }
      ]
    }
  },
  {
    id: 'func-3',
    name: 'formatCurrency',
    visibility: 'public',
    description: 'Formats a number as currency with locale support',
    position: { x: 100, y: 400 },
    sourceLocation: {
      file: 'src/utils/format.js',
      startLine: 8,
      endLine: 12
    },
    signature: {
      name: 'formatCurrency',
      parameters: [
        {
          name: 'amount',
          type: 'number',
          description: 'Amount to format',
          optional: false
        },
        {
          name: 'locale',
          type: 'string',
          description: 'Locale code',
          optional: true,
          defaultValue: 'en-US'
        },
        {
          name: 'currency',
          type: 'string',
          description: 'Currency code',
          optional: true,
          defaultValue: 'USD'
        }
      ],
      returns: [
        {
          type: 'string',
          description: 'Formatted currency string'
        }
      ]
    }
  }
];
