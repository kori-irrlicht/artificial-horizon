package main

import (
	"fmt"
	_ "image/png"
	"io/ioutil"
	"log"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/kori-irrlicht/artificial-horizon/core"
)

var _ core.Game = &game{}

type game struct {
	window     *glfw.Window
	controller core.Controller
	vao        uint32
	program    uint32
	texture    uint32
	modelUni   int32
	model      mgl32.Mat4
	angle      float64
	prev       float64
}

func (g *game) Update() {}
func (g *game) Input() {
	if g.controller.IsDown(KeyDown) {
		fmt.Println("Down")
	}
	if g.controller.IsDown(KeyUp) {
		fmt.Println("Up")
	}
}
func (g *game) Render() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	fmt.Println(g)
	time := glfw.GetTime()
	elapsed := time - g.prev
	g.prev = time

	g.angle += elapsed

	gl.UseProgram(g.program)
	gl.UniformMatrix4fv(g.modelUni, 1, false, &g.model[0])
	gl.BindVertexArray(g.vao)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, g.texture)

	gl.DrawArrays(gl.TRIANGLES, 0, 2*2*3)

	g.window.SwapBuffers()
	glfw.PollEvents()
}
func (g *game) FrameTime() time.Duration {
	return 16 * time.Millisecond
}

func (g *game) Now() time.Time {
	return time.Now()
}

func (g *game) Running() bool {
	return !g.window.ShouldClose()
}

func newGame() (core.Game, error) {
	game := &game{}

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(vidMode.Width, vidMode.Height, "Artificial Horizon", glfw.GetPrimaryMonitor(), nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	game.window = window

	kcm, _ := core.NewKeyCallbackManager(window)
	mapping := core.KeyboardMapping{
		{KeyUp, glfw.KeyW},
		{KeyDown, glfw.KeyS},
		{KeyLeft, glfw.KeyA},
		{KeyRight, glfw.KeyD},
	}
	game.controller, _ = core.NewKeyboardController(kcm, mapping)

	//`conn, err := network.NewConnection("127.0.0.1", "42425", "42426")
	if err != nil {
		panic(err)
	}

	// conn.Tcp().Write([]byte("Hallo"))

	cubeVertices := []float32{
		//  X, Y, Z, U, V

		// Front
		-1.0, -1.0, 1.0, 1.0, 0.0,
		1.0, -1.0, 1.0, 0.0, 0.0,
		-1.0, 1.0, 1.0, 1.0, 1.0,
		1.0, -1.0, 1.0, 0.0, 0.0,
		1.0, 1.0, 1.0, 0.0, 1.0,
		-1.0, 1.0, 1.0, 1.0, 1.0,

		// X Y Z U V
		-3.0, -1.0, 1.0, 1.0, 0.0,
		-1.0, -1.0, 1.0, 0.0, 0.0,
		-3.0, 1.0, 1.0, 1.0, 1.0,
		-1.0, -1.0, 1.0, 0.0, 0.0,
		-1.0, 1.0, 1.0, 0.0, 1.0,
		-3.0, 1.0, 1.0, 1.0, 1.0,
	}

	vertexShader, err := ioutil.ReadFile("./shader/vertex.glsl")
	fragmentShader, err := ioutil.ReadFile("./shader/fragment.glsl")

	// Configure the vertex and fragment shaders
	program, err := newProgram(string(vertexShader)+"\x00", string(fragmentShader)+"\x00")
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	//projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(vidMode.Width)/float32(vidMode.Height), 0.1, 10.0)
	projection := mgl32.Ortho(-5, 5, -5, 5, 0.1, 10)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	camera := mgl32.LookAtV(mgl32.Vec3{0, 7, 6}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})

	//camera := mgl32.LookAtV(mgl32.Vec3{0, 0, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 0, 0})
	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &camera[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	// Load the texture
	game.texture, err = newTexture("./assets/square.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(cubeVertices)*4, gl.Ptr(cubeVertices), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointer(vertAttrib, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))

	texCoordAttrib := uint32(gl.GetAttribLocation(program, gl.Str("vertTexCoord\x00")))
	gl.EnableVertexAttribArray(texCoordAttrib)
	gl.VertexAttribPointer(texCoordAttrib, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))

	// Configure global settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	game.program = program
	game.vao = vao
	game.model = model
	game.modelUni = modelUniform

	fmt.Println(game)

	return game, nil
}
