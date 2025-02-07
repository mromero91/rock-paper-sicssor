
# ✊🖐✌ Rock-Paper-Scissor (Piedra, Papel o Tijera)

Un clásico juego de **Piedra, Papel o Tijera** desarrollado en **Go** con una interfaz simple y atractiva.

## 📦 Clonar el Proyecto

```sh
git clone https://github.com/tuusuario/rock-paper-scissor.git
cd rock-paper-scissor
```
⚙️ Instalación de Dependencias

go mod tidy

🚀 Configuración para Desarrollo

Para facilitar el desarrollo con hot-reload, puedes instalar Air:

go install github.com/cosmtrek/air@latest

🌱 Modo Desarrollo (con Air)

Ejecuta el siguiente comando para iniciar el proyecto con recarga en vivo:

air

🚀 Modo Producción

Compila y ejecuta el proyecto manualmente:

go build -o tmp/main cmd/main.go
./tmp/main

📁 Estructura del Proyecto

rock-paper-scissor/
├── cmd/                # Código principal de la aplicación
│   └── main.go
├── templates/          # Archivos HTML para la interfaz
│   ├── base.html
│   ├── index.html
│   ├── game.html
│   ├── new-game.html
│   └── about.html
├── static/             # Archivos estáticos como imágenes y scripts
│   ├── js/
│   │   └── main.js
│   └── img/
├── tmp/                # Archivos temporales y compilaciones
├── .air.toml           # Configuración de Air
├── go.mod              # Archivo de módulos de Go
└── README.md           # Este archivo 😃

🛠 Tecnologías Utilizadas
	•	Go 🦫
	•	Tailwind CSS 🎨
	•	HTML5 + JavaScript 🖥️
	•	Air (hot-reload) 🔄

🤝 Contribuir

¡Las contribuciones son bienvenidas! Para contribuir:
	1.	Haz un fork del repositorio.
	2.	Crea una nueva rama (git checkout -b feature-nueva-funcionalidad).
	3.	Realiza los cambios y haz commit (git commit -m "Agrega nueva funcionalidad").
	4.	Envía un pull request.

📄 Licencia

Este proyecto está bajo la licencia MIT.

💡 Diviértete jugando y programando! 🚀

---

### **🔹 Mejoras y cambios en este README:**
✅ **Estructura clara** con títulos y secciones bien definidas.  
✅ **Uso de emojis** para mejorar la legibilidad y atractivo visual.  
✅ **Explicaciones detalladas** para la instalación y ejecución.  
✅ **Añadida la estructura del proyecto** para mayor claridad.  
✅ **Sección de contribución** para invitar a mejoras y colaboraciones.  

📌 **¡Ahora tu README se ve profesional y atractivo en GitHub!** 🚀💻


