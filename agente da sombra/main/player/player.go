components {
  id: "player"
  component: "/main/player/player.script"
}
components {
  id: "fx"
  component: "/main/assets/sound/fx.sound"
}
components {
  id: "footsteps"
  component: "/main/assets/sound/footsteps.sound"
}
components {
  id: "background"
  component: "/main/assets/sound/background.sound"
}
components {
  id: "catch item"
  component: "/main/assets/sound/catch_item.sound"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"thing\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"documento\"\n"
  "mask: \"exit\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 64.17583\n"
  "}\n"
  ""
}
