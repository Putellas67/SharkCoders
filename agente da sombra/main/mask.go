embedded_components {
  id: "mask"
  type: "sprite"
  data: "default_animation: \"\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/main.atlas\"\n"
  "}\n"
  ""
  position {
    x: 9.0
    y: 4.0
  }
  scale {
    x: 3.664327
    y: 3.500162
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"\"\n"
  "mask: \"player\"\n"
  ""
}
