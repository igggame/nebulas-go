// Copyright (C) 2018 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or
// modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see
// <http://www.gnu.org/licenses/>.
//
#pragma once
#include "compatible/compatible_check_interface.h"

namespace neb {
namespace compatible {

class compatible_check_base : public compatible_check_interface {
public:
  compatible_check_base(version_t v);
  virtual ~compatible_check_base();

  virtual version_t rt_version() const;

protected:
  version_t m_version;
};
} // namespace compatible
} // namespace neb
