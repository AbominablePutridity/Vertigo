components {
  id: "player_sprite"
  component: "/main/prefabs/platform/platform_sprite.sprite"
  position {
    x: -0.00547395
    y: 0.0018895074
  }
  scale {
    x: 2.0
    y: 2.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.0\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "mask: \"knife\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 60.0\n"
  "  data: 30.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
