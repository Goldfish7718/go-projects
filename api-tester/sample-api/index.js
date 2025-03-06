import express from 'express';

const app = express();

app.use(express.json());

app.get('/', (_, res) => res.send('Meow'));
app.get('/json', (_, res) => res.json({ message: 'Now this some json right here' }));

app.post('/', (req, res) => {
  const { body } = req;
  console.log(req);

  res.json({ body });
});
app.put('/', (req, res) => res.send('put'));
app.delete('/', (req, res) => res.send('Delete'));
app.listen(3000, () => console.log('renning'));
