import express from 'express';

const app = express();
const jsonResponse = {
  first: 'Kendrick Lamar',
  second: 'Drake',
  third: 'J cole',
  array: ['TRAVIS SCOTT', 'YOUNG THUG'],
};

app.use(express.json());

app.get('/', (_, res) => res.send('Meow'));
app.get('/json', (_, res) => res.json(jsonResponse));

app.post('/', (req, res) => {
  const { body } = req;

  res.json({ ...body });
});

app.put('/', (req, res) => {
  const { body } = req;

  res.json({ ...body });
});

app.delete('/', (req, res) => res.send('Delete'));
app.listen(3000, () => console.log('renning'));
