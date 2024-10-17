import fs from 'fs'
import path from 'path'
import { Hono } from 'hono'
import { serve } from '@hono/node-server'
import { serveStatic } from '@hono/node-server/serve-static'

function readAsset(Path: string) {
	const filePath = path.join('src', 'build', Path)
	const fileContent = fs.readFileSync(filePath, 'utf-8')
	return fileContent
}

const app = new Hono()

app.use('/assets/*', serveStatic({ root: './src/build/' }))

app.use('/images/*', serveStatic({ root: './src/build/' }))

app.get('/', (c) => {
	return c.html(readAsset('index.html'))
})

app.get('/searching', (c) => {
	return c.html(readAsset('searching.html'))
})

app.get('/blog/:slug', (c) => {
	const { slug } = c.req.param()
	return c.html(readAsset(`blog/${slug}.html`))
})

const port = 3000
console.log(`Server is running on port ${port}`)

serve({
	fetch: app.fetch,
	port,
})
