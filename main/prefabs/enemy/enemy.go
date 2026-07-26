components {
  id: "enemy_sprite"
  component: "/main/prefabs/enemy/enemy.sprite"
  position {
    x: 0.007142857
    y: 0.025
  }
}
components {
  id: "enemy_controller"
  component: "/main/prefabs/enemy/scripts/enemy_controller.script"
}
embedded_components {
  id: "enem"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.0\n"
  "group: \"enemy\"\n"
  "mask: \"default\"\n"
  "mask: \"knife\"\n"
  "mask: \"patrol_wall\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 13.0\n"
  "  data: 43.5\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Pistol\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 40.0\n"
  "  y: 20.0\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/images/Pistol.atlas\"\n"
  "}\n"
  ""
  position {
    x: -20.0
    y: -16.0
    z: 7.0
  }
}
