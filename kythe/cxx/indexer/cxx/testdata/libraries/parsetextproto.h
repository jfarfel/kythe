#ifndef KYTHE_CXX_INDEXER_CXX_TESTDATA_LIBRARIES_PROTO_PARSETEXTPROTO_H_
#define KYTHE_CXX_INDEXER_CXX_TESTDATA_LIBRARIES_PROTO_PARSETEXTPROTO_H_

#include "absl/strings/string_view.h"
#include "absl/types/source_location.h"

// The following declarations mimick the code structure of ParseProtoHelper.
struct StringPiece {
  StringPiece(const char* str) {}  // NOLINT
};

namespace proto2 {
class Message;
namespace contrib {
namespace parse_proto {
namespace internal {
struct ParseProtoHelper {
  template <typename T>
  operator T() const;
};
}  // namespace internal

internal::ParseProtoHelper ParseTextProtoOrDie(
    absl::string_view asciipb, ParserConfig config,
    absl::SourceLocation loc ABSL_LOC_CURRENT_DEFAULT_ARG);

}  // namespace parse_proto
}  // namespace contrib
}  // namespace proto2

#define PARSE_TEXT_PROTO(asciipb) \
    ::proto2::contrib::parse_proto::ParseTextProtoOrDie(asciipb, {}, ABSL_LOC)

#endif  // KYTHE_CXX_INDEXER_CXX_TESTDATA_LIBRARIES_PROTO_PARSETEXTPROTO_H_
