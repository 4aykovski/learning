using Microsoft.EntityFrameworkCore;
using TaskManager.Model;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace TaskManager.Service.EntityTypeConfiguration;

public class NoteConfiguration : IEntityTypeConfiguration<Note>
{
    public void Configure(EntityTypeBuilder<Note> builder)
    {
        builder.HasKey(note => note.Id);
        builder.HasIndex(note => note.Id).IsUnique();
        builder.Property(note => note.Title).HasMaxLength(100).IsRequired();
    }
}