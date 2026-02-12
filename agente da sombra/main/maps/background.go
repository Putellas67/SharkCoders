components {
  id: "background"
  component: "/main/assets/sound/background.sound"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"background\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/main.atlas\"\n"
  "}\n"
  ""
}
