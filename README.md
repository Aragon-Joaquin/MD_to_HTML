Updating when i have an idea how to make a decent start with this

My main idea was to implement some type of Abstract Syntax Tree to keep track of Symbols like \*, \_, $, = and ```.

### So what im gonna do is:

These steps are the same as a _compiler_ uses. It's **overkill** for a simple processor like this but it there's so much juice to learn while doing it like this.

_SPANISH ALERT ❗❗❗❗_

1. **_Tokenaizer_ / Lexer**

   La idea de esta fase es transformar lineas de markdown en la unidad mas pequeña posible, en este caso, simbolos y palabras individuales para poder destacar la sintaxis y armar el AST. A este output lo llamariamos "Token", y se veria algo asi:

   ```js
   // example.md
   // **Bold Text**
   const Lexer = [
   	{ type: 'Symbol', value: '*' },
   	{ type: 'Symbol', value: '*' },
   	{ type: 'String', value: 'Bold' },
   	{ type: 'String', value: 'text' },
   	{ type: 'Symbol', value: '*' },
   	{ type: 'Symbol', value: '*' }
   ]
   ```

| **Token Type** | **Elements**                                                              |
| -------------- | ------------------------------------------------------------------------- |
| Symbol         | `*`, `_`, `$`, `=`, ` `` `, `~`, `#`, `>`, `-`, `[ and ]`, `( and )`, `!` |
| String         | Words, phrases, or plain text                                             |

2. **_Parser_**

   Es aqui donde se define el AST (_Abstract Syntax Tree_ ). Aqui es donde los tokens se van a transformar en nodos.
   Como ya explique anteriormente, el AST va resultar bastante breve donde los parents nodes simplemente serán los **Symbols**.
   He aqui un ejemplo de lo que tengo en mente:

   ```js
   const parsed = [
   	{
   		type: 'Symbol',
   		value: '**',
   		repeatsAtTheEnd: true,
   		body: [
   			{
   				type: 'String',
   				value: 'Bold',
   				repeatsAtTheEnd: false,
   				body: []
   			},
   			{
   				type: 'String',
   				value: 'text',
   				repeatsAtTheEnd: false,
   				body: []
   			}
   		]
   	},
   	// and if we would have another token...
   	{
   		type: 'String',
   		value: 'hello',
   		repeatsAtTheEnd: false,
   		body: []
   	}
   ]
   ```

3. **_Traverser_ (not common on most compilers)**

   La idea de este punto era simplemente realizar un hashmap de funciones abstractas e injectarlas al nodo simplemente para agilizar el recorrido de sus hijos haciendolo recursivamente, no creo implementarlo.

4. **_Transformer_**

   Para este ejemplo, el transformer no tiene mucho sentido emplearlo ya que se suele hacer optimizaciones y añadir mas informacion tecnica de la syntaxis. Pero aun asi lo utilizaremos para designar el tipo de expresion que es, y cual es la etiqueta HTML que le pertenece.

5. **_HTML Output_**

   .
