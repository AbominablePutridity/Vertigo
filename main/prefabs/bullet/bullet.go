components {
  id: "physic_throwing"
  component: "/main/prefabs/physic_throwing.script"
  properties {
    id: "speed"
    value: "500.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "flight_weapon"
  component: "/main/prefabs/flight_weapon.script"
  properties {
    id: "is_knife"
    value: "false"
    type: PROPERTY_TYPE_BOOLEAN
  }
}
embedded_components {
  id: "bullet_sprite"
  type: "sprite"
  data: "default_animation: \"Bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 25.0\n"
  "  y: 12.5\n"
  "}\n"
  "size_mode: SIZE_MODE_MANUAL\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/images/Bullet.atlas\"\n"
  "}\n"
  ""
  rotation {
    y: 1.0
    w: 6.123234E-17
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 0.5\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
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
  "  data: 11.0\n"
  "  data: 5.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
