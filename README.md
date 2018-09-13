Crud REST API Go

Use o config.toml.example para criar a configuração de seu MongoDB

Acessos

GET    http://localhost:3000/api/v1/movies
GET    http://localhost:3000/api/v1/movies/{id}
POST   http://localhost:3000/api/v1/movies
PUT    http://localhost:3000/api/v1/movies/{id}
DELETE http://localhost:3000/api/v1/movies/{id}

Exemplos:

Create
{
  "name": "The Equalizer",
  "description": "Robert McCall (Denzel Washington), a man of mysterious origin who believes he has put the past behind him, dedicates himself to creating a quiet new life. However, when he meets Teri (Chloë Grace Moretz), a teenager who has been manhandled by violent Russian mobsters, he simply cannot walk away. With his set of formidable skills, McCall comes out of self-imposed retirement and emerges as an avenging angel, ready to take down anyone who brutalizes the helpless.",
  "thumb_image": "http://t2.gstatic.com/images?q=tbn:ANd9GcQkGfxoavBEp4fR6P-yi2mIkUl1aZHHFIietLK4GriI5YyvGSJ7",
  "active": true
}

Update
{
	"id":"5b8ab0132e9a2421907d5669",
  "name": "The Equalizer 2",
  "description": "If you have a problem and there is nowhere else to turn, the mysterious and elusive Robert McCall will deliver the vigilante justice you seek. This time, however, McCall's past cuts especially close to home when thugs kill Susan Plummer -- his best friend and former colleague. Now out for revenge, McCall must take on a crew of highly trained assassins who'll stop at nothing to destroy him.",
  "thumb_image": "http://t1.gstatic.com/images?q=tbn:ANd9GcRCss5jFvU87fla4jEIwpv9dUAdZKzeUYHY_mhqzDxrvX9ppWyJ",
  "active": true
}

